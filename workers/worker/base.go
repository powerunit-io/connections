package worker

import (
	"fmt"

	"github.com/powerunit-io/platform/config"
	"github.com/powerunit-io/platform/logging"
)

// BaseWorker -
type BaseWorker struct {
	*logging.Logger
	*config.Config
}

// Validate -
func (bw *BaseWorker) Validate() error {
	bw.Debug("Validating basic configurations for (worker: %s)", bw.String())

	if _, ok := bw.Config.Get("worker_name").(string); !ok {
		return fmt.Errorf(
			"Could not build worker as base worker name is missing or is not valid (entry: %s)",
			bw.Config.Get("worker_name"),
		)
	}

	if len(bw.Config.Get("worker_name").(string)) < 2 {
		return fmt.Errorf(
			"Could not build worker as base worker name is way to short. It needs to be at least 2 chars long. (entry: %s)",
			bw.Config.Get("worker_name"),
		)
	}

	return nil
}

// String - We use it to name worker when needed.
func (bw BaseWorker) String() string {
	if worker, ok := bw.Config.Get("worker_name").(string); ok {
		return worker
	}

	return ""
}
