package mysql

import (
	"github.com/powerunit-io/platform/config"
	"github.com/powerunit-io/platform/logging"
)

// Manager -
type Manager interface {
	Start(done chan bool) chan error
	Validate() error
	WorkerName() string
	Stop() error
}

// NewConnection -
func NewConnection(log *logging.Logger, conf *config.Config) (Manager, error) {
	return Manager(&Connection{
		Logger: log, Config: conf, Uri: conf.Get("uri").(string),
	}), nil
}
