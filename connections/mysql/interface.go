package mysql

import (
	"database/sql"

	"github.com/powerunit-io/platform/config"
	"github.com/powerunit-io/platform/logging"
)

// Manager -
type Manager interface {
	Name() string

	Validate() error

	Start(done chan bool) chan error
	Stop() error

	Conn() *sql.DB
}

// NewConnection -
func NewConnection(log *logging.Logger, conf *config.Config) (Manager, error) {
	return Manager(&Connection{
		Logger: log, Config: conf, Uri: conf.Get("uri").(string),
	}), nil
}
