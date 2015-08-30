package mqtt

import (
	"github.com/powerunit-io/platform/config"
	"github.com/powerunit-io/platform/logging"
	"github.com/powerunit-io/platform/workers/worker"
)

// MqttConnection -
type MqttConnection struct {
	*logging.Logger
	*config.Config
	worker.BaseWorker
}

// Start -
func (mc *MqttConnection) Start(done chan bool) error {
	mc.Warning("Starting mqtt device (connection: %s) ...", "")
	return nil
}

// Validate -
func (mc *MqttConnection) Validate() error {
	return nil
}

// Stop -
func (mc *MqttConnection) Stop() error {
	mc.Warning("Stopping mqtt device (connection: %s) ...", "")
	return nil
}
