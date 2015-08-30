package mqtt

import (
	"fmt"

	"github.com/powerunit-io/platform/config"
	"github.com/powerunit-io/platform/logging"
	"github.com/powerunit-io/platform/workers/worker"

	MQTT "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
)

// MqttConnection -
type MqttConnection struct {
	*logging.Logger
	*config.Config
	*worker.BaseWorker

	conn *MQTT.Client
}

// Start -
func (mc *MqttConnection) Start(done chan bool) error {
	mc.Warning("Starting mqtt (worker: %s) on (addr: %s)...", mc.Config.Get("worker_name"), mc.GetBrokerAddr())

	opts := MQTT.NewClientOptions().AddBroker(mc.GetBrokerAddr())
	opts.SetClientID(mc.GetBrokerClientId())
	opts.SetDefaultPublishHandler(mc.BrokerHandler)

	mc.conn = MQTT.NewClient(opts)

	if token := mc.conn.Connect(); token.Wait() && token.Error() != nil {
		return fmt.Errorf(
			"Failed to establish connection with mqtt server (error: %s)",
			token.Error(),
		)
	}

	return nil
}

// BrokerHandler -
func (mc *MqttConnection) BrokerHandler(client *MQTT.Client, msg MQTT.Message) {

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

// GetBrokerAddr -
func (mc *MqttConnection) GetBrokerAddr() string {
	connection := mc.Config.Get("connection").(map[string]interface{})
	return fmt.Sprintf("%s://%s", connection["network"].(string), connection["address"].(string))
}

// GetBrokerClientId -
func (mc *MqttConnection) GetBrokerClientId() string {
	connection := mc.Config.Get("connection").(map[string]interface{})
	return connection["clientId"].(string)
}

// Stop -
func (mc *MqttConnection) Stop() error {
	mc.Warning("Stopping mqtt device (connection: %s) ...", "")
	return nil
}
