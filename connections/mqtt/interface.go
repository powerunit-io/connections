package mqtt

import (
	"github.com/powerunit-io/platform/config"
	"github.com/powerunit-io/platform/logging"
)

// Manager -
type Manager interface {
	Start(done chan bool) chan error
	Validate() error
	String() string
	Stop() error
}

// NewConnection -
func NewConnection(log *logging.Logger, conf *config.Config) (Manager, error) {
	return Manager(&Connection{Logger: log, Config: conf}), nil
}
