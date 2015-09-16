// Copyright 2015 The PowerUnit Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package managers ...
package managers

// Service -
type Service interface {
	Start(done chan bool) error
	Stop() error
	Validate() error
	Name() string

	// Adapter - Will return instance of service. If it's mysql, it will
	// return mysql.Connection instance
	// remember that you will need to cast it to proper type before it can be used
	// example: db = service.Adapter().(*mysql.Connection) fmt.Printf("%q", db.DB)
	Adapter() interface{}
}

// Manager -
type Manager interface {
	Attach(m string, bm Service) error
	Remove(m string) error
	All() map[string]Service
	List() []string
	Get(m string) (Service, error)
	Exists(m string) bool
}
