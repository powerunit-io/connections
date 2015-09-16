// Copyright 2015 The PowerUnit Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package mysql ...
package mysql

import (
	"database/sql"
	"fmt"
	"time"

	//
	_ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"
	"github.com/powerunit-io/platform/config"
	"github.com/powerunit-io/platform/logging"
	"github.com/powerunit-io/platform/utils"
)

// Connection -
type Connection struct {
	*logging.Logger
	*config.Config
	URI string
	gorm.DB
}

// Start - Will connect to database and than try to reconnect in case that we get disconnected
func (m *Connection) Start(done chan bool) error {
	m.Info("Starting MSQL Connection: %s ...", m.URI)

	// Keep track of first connection, blocking, reconnect in background
	started := make(chan bool)
	connected := false
	reconnect := make(chan bool)

	// Reconnect background loop
	go func() {
		for {
			// Try to connect
			err := m.Connect()

			if err != nil {
				m.Logger.Error("Got Error while connecting against MYSQL: %s", err)
				time.Sleep(5 * time.Second)
				continue
			}

			// Notify of first connection
			if !connected {
				connected = true
				started <- connected
			}

			go func() {
				for {
					if !m.IsConnected() {
						m.Logger.Error("Could not connect to MySQL server. Reconnecting in 5 seconds ...")
						time.Sleep(5 * time.Second)
						reconnect <- true
						continue
					}

					time.Sleep(5 * time.Second)
				}
			}()

			<-reconnect
		}
	}()

	// Wait for background loop to connect for the first time
	<-started
	m.Info("MYSQL Connection successfully established: %v", m)

	return nil
}

// Validate -
func (m *Connection) Validate() error {
	m.Info("Validating mysql configuration for (name: %s)", m.Name())

	// @TODO - This needs proper regex validation ...
	if len(m.URI) < 10 {
		return fmt.Errorf(
			"Failed to validate mysql connection (name: %s) uri. You've passed (uri: %s)",
			m.Name(), m.URI,
		)
	}

	return nil
}

// Connect - Will connect and ping connection in hope that all is ok
func (m *Connection) Connect() error {
	m.Logger.Debug("Connecting to MySQL server (uri: %s) ...", m.URI)

	var err error

	m.DB, err = gorm.Open("mysql", m.URI)
	//m.DB.SetLogger(m.Logger)

	if err != nil {
		m.Logger.Error("Got error while attempting to create mysql connection: %s", err)
		return err
	}

	if err := m.DB.DB().Ping(); err != nil {
		m.Logger.Error("Got error while attempting to PING database: %s", err)
		return err
	}

	concurrency := utils.GetConcurrencyCount("PU_GO_MAX_CONCURRENCY")

	// Setting up max idle conns based on concurrency
	m.DB.DB().SetMaxIdleConns(int(concurrency))
	m.DB.DB().SetMaxOpenConns(int(concurrency))

	return nil
}

// Commit - Will basically prepare query and execute it returning back both error (if any) and result
func (m *Connection) Commit(query string, args ...interface{}) (sql.Result, error) {
	stmt, err := m.DB.DB().Prepare(query)

	if err != nil {
		m.Logger.Error("Error happen while attempting to prepare query: %s", err)
		return nil, err
	}

	defer stmt.Close()

	m.Logger.Info("About to execute following mysql (query: %s) - (args: %v)", query, args)

	return stmt.Exec(args...)
}

// IsConnected - Check whenever we can ping database or not
func (m *Connection) IsConnected() bool {
	if err := m.DB.DB().Ping(); err != nil {
		m.Logger.Error("Got error while attempting to PING database: %s", err)
		return false
	}

	return true
}

// Adapter -
func (m *Connection) Adapter() interface{} {
	return m
}

// Name -
func (m *Connection) Name() string {
	return m.Config.Get("name").(string)
}

// Stop - Will close MySQL connection if we ever need it
func (m *Connection) Stop() error {
	m.Warning(
		"Closing MySQL connection for (name: %s) - (uri: %s) ...",
		m.Name(), m.URI,
	)

	return m.Close()
}
