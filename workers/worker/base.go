package worker

import (
	"github.com/powerunit-io/platform/config"
	"github.com/powerunit-io/platform/logging"
)

// BaseWorker -
type BaseWorker struct {
	*logging.Logger

	Config *config.ConfigManager
}

// Validate -
func (bw *BaseWorker) Validate() error {

	return nil
}

// Start -
func (bw *BaseWorker) Start(done chan bool) error {

	return nil
}
