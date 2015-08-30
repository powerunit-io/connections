package mqtt

import (
	"fmt"

	"github.com/powerunit-io/platform/config"
	"github.com/powerunit-io/platform/logging"
	"github.com/powerunit-io/platform/workers/worker"
)

// MqttConnection -
type MqttConnection struct {
	*logging.Logger
	*config.Config
	*worker.BaseWorker
}

// Start -
func (mc *MqttConnection) Start(done chan bool) error {
	mc.Warning("Starting mqtt (worker: %s) ...", mc.Config.Get("worker_name"))
	return nil
}

// Validate -
func (mc *MqttConnection) Validate() error {
	mc.Debug("Validating mqtt configuration for (worker: %q)", mc.Config.Get("worker_name"))

	if mc.Config.Get("connection") == nil {
		return fmt.Errorf(
			"Could not build mqtt worker as connection interface is missing (entry: %s)",
			mc.Config.Get("connection"),
		)
	}

	data := mc.Config.Get("connection").(map[string]interface{})

	if _, ok := data["network"].(string); !ok {
		return fmt.Errorf(
			"Could not mqtt worker as connection network is not set. (connection_data: %q)",
			data,
		)
	}

	if _, ok := data["address"].(string); !ok {
		return fmt.Errorf(
			"Could not mqtt worker as connection address is not set. (connection_data: %q)",
			data,
		)
	}

	if _, ok := data["clientId"].(string); !ok {
		return fmt.Errorf(
			"Could not mqtt worker as connection clientId is not set. (connection_data: %q)",
			data,
		)
	}

	return nil
}

// Stop -
func (mc *MqttConnection) Stop() error {
	mc.Warning("Stopping mqtt device (connection: %s) ...", "")
	return nil
}
