package mqtt

import (
	"fmt"
	"strings"
	"time"

	"github.com/powerunit-io/platform/config"
	"github.com/powerunit-io/platform/events"
	"github.com/powerunit-io/platform/logging"
	"github.com/powerunit-io/platform/utils"
	"github.com/powerunit-io/platform/workers/worker"

	MQTT "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
)

// MqttConnection -
type MqttConnection struct {
	*logging.Logger
	*config.Config
	*worker.BaseWorker

	conn   *MQTT.Client
	events chan events.Event
}

// Start -
func (mc *MqttConnection) Start(done chan bool) chan error {
	opts := MQTT.NewClientOptions().AddBroker(mc.GetBrokerAddr())
	opts.SetClientID(mc.GetBrokerClientId())
	opts.SetDefaultPublishHandler(mc.BrokerHandler)
	opts.SetAutoReconnect(true)

	errors := make(chan error)
	connected := make(chan bool)

	go func() {
		for {
			mc.Warning("Starting mqtt (worker: %s) on (addr: %s)...", mc.WorkerName(), mc.GetBrokerAddr())

			reload := make(chan bool)
			mc.conn = MQTT.NewClient(opts)

			if token := mc.conn.Connect(); token.Wait() && token.Error() != nil {
				errors <- fmt.Errorf(
					"Failed to establish connection with mqtt server (error: %s)",
					token.Error(),
				)
				continue
			}

			// In case we hit connected, ensure that main func is notified
			if !mc.conn.IsConnected() {
				continue
			}

			mc.Subscribe(mc.GetBrokerTopicName(), MaxTopicSubscribeAttempts)

			// Notify rest of the app that we're ready ...
			close(connected)

			go func() {
				cct := time.Tick(2 * time.Second)

				for {
					select {
					case <-cct:
						if !mc.conn.IsConnected() {
							reload <- true
							return
						}
					case <-done:
						mc.Warning("Received stop signal for mqtt (worker: %s). Will not attempt to restart worker ...", mc.WorkerName())
						return
					}
				}
			}()

		reloadloop:
			for {
				select {
				case <-reload:
					mc.Warning("Mqtt (worker: %s) seems not to be connected. Restarting loop in 2 seconds ...", mc.WorkerName())
					time.Sleep(2 * time.Second)
					break reloadloop
				}
			}

		}
	}()

	select {
	case <-connected:
		mc.Info(
			"Successfully established mqtt connection for (worker: %s) on (addr: %s)",
			mc.WorkerName(), mc.GetBrokerAddr(),
		)
		break
	case <-time.After(time.Duration(InitialConnectionTimeout) * time.Second):
		errors <- fmt.Errorf(
			"Could not establish mqtt connection for (worker: %s) on (addr: %s) due to initial connection (timeout: %ds)",
			mc.WorkerName(), mc.GetBrokerAddr(), InitialConnectionTimeout,
		)
		break
	}

	return errors
}

// GetEvents -
func (mc *MqttConnection) DrainEvents() chan events.Event {
	return mc.events
}

// Subscribe -
func (mc *MqttConnection) Subscribe(topic string, maxRetryAttempts int) error {
	var err error

	for i := 0; i <= maxRetryAttempts; i++ {
		mc.Debug(
			"About to attempt subscribe to mqtt (topic: %s) for (worker: %s) -> (retry_attempt: %d)",
			topic, mc.WorkerName(), i,
		)

		if token := mc.conn.Subscribe(topic, 0, nil); token.Wait() && token.Error() != nil {
			mc.Error("Could not subscribe to (topic: %s) for (worker: %s) due to (err: %s). Retrying ...")
			err = token.Error()
			continue
		}

		mc.Info("Successfully subscribed (worker: %s) on (topic: %s)!",
			mc.WorkerName(), mc.GetBrokerTopicName(),
		)

		err = nil
		break
	}

	return err
}

// BrokerHandler -
func (mc *MqttConnection) BrokerHandler(client *MQTT.Client, msg MQTT.Message) {
	mc.Debug(
		"Received new mqtt (worker: %s) - (message: %s) for (topic: %s). Building event now ...",
		mc.WorkerName(), msg.Payload(), msg.Topic(),
	)

	event := events.NewEvent(client, msg)
	mc.Debug("%v", event)
}

// Validate -
func (mc *MqttConnection) Validate() error {
	mc.Debug("Validating mqtt configuration for (worker: %q)", mc.WorkerName())

	if mc.Config.Get("connection") == nil {
		return fmt.Errorf(
			"Could not validate mqtt worker as connection interface is missing (entry: %s)",
			mc.Config.Get("connection"),
		)
	}

	data := mc.Config.Get("connection").(map[string]interface{})

	if _, ok := data["network"].(string); !ok {
		return fmt.Errorf(
			"Could not validate mqtt worker as connection network is not set. (connection_data: %q)",
			data,
		)
	}

	if !utils.StringInSlice(data["network"].(string), AvailableConnectionTypes) {
		return fmt.Errorf(
			"Could not validate mqtt worker as connection network is not valid. (network: %s) - (available_networks: %v)",
			data["network"].(string), AvailableConnectionTypes,
		)
	}

	if _, ok := data["address"].(string); !ok {
		return fmt.Errorf(
			"Could not validate mqtt worker as connection address is not set. (connection_data: %q)",
			data,
		)
	}

	address := data["address"].(string)

	if len(address) < 5 || !strings.Contains(address, ":") {
		return fmt.Errorf(
			"Could not validate mqtt worker as connection address is not valid. (address: %s)",
			address,
		)
	}

	clientId := data["clientId"].(string)

	if _, ok := data["clientId"].(string); !ok {
		return fmt.Errorf(
			"Could not validate mqtt worker as connection clientId is not set. (client_id: %s)",
			clientId,
		)
	}

	if len(clientId) < 2 {
		return fmt.Errorf(
			"Could not validate mqtt worker as connection clientId is not long enough. (client_id: %s)",
			clientId,
		)
	}

	if _, ok := data["topic"].(string); !ok {
		return fmt.Errorf(
			"Could not validate mqtt worker as connection topic is not set. (connection_data: %q)",
			data,
		)
	}

	return nil
}

// GetBrokerAddr -
func (mc *MqttConnection) GetBrokerAddr() string {
	connection := mc.Config.Get("connection").(map[string]interface{})
	return fmt.Sprintf("%s://%s?timeout=10s", connection["network"].(string), connection["address"].(string))
}

// GetBrokerClientId -
func (mc *MqttConnection) GetBrokerClientId() string {
	connection := mc.Config.Get("connection").(map[string]interface{})
	return connection["clientId"].(string)
}

// GetBrokerTopicName -
func (mc *MqttConnection) GetBrokerTopicName() string {
	connection := mc.Config.Get("connection").(map[string]interface{})
	return connection["topic"].(string)
}

// WorkerName -
func (mc *MqttConnection) WorkerName() string {
	return mc.Config.Get("worker_name").(string)
}

// Stop -
func (mc *MqttConnection) Stop() error {
	mc.Warning("Stopping mqtt (worker: %s) ...", mc.WorkerName())

	if !mc.conn.IsConnected() {
		mc.Warning("Connection for mqtt (worker: %s) is already closed.", mc.WorkerName())
		return nil
	}

	mc.Warning("Unsubscribing from mqtt (worker: %s) (topic: %s)...", mc.WorkerName(), mc.GetBrokerTopicName())
	if token := mc.conn.Unsubscribe(mc.GetBrokerTopicName()); token.Wait() && token.Error() != nil {
		mc.Error(
			"Could not unsubscribe from (topic: %s) for (worker: %s) due to (err: %s)",
			mc.GetBrokerTopicName(), mc.WorkerName(), token.Error(),
		)
	}

	mc.Warning("Stopping mqtt (worker: %s) connection...", mc.WorkerName())
	mc.conn.Disconnect(2)
	time.Sleep(2 * time.Second)

	return nil
}
