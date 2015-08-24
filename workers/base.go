package workers

import (
	"github.com/powerunit-io/platform/config"
)

type BaseWorker struct {
	Service interface{}
	Config  *config.ConfigManager
}
