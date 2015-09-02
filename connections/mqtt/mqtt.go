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

// Connection -
type Connection struct {
	*logging.Logger
	*config.Config
	*worker.BaseWorker

	conn   *MQTT.Client
	events chan events.Event
}

// Start -
func (c *Connection) Start(done chan bool) chan error {
	opts := MQTT.NewClientOptions().AddBroker(c.GetBrokerAddr())
	opts.SetClientID(c.GetBrokerClientId())
	opts.SetDefaultPublishHandler(c.BrokerHandler)
	opts.SetAutoReconnect(true)

	concurrency := utils.GetConcurrencyCount("PU_GO_MAX_CONCURRENCY")
	c.events = make(chan events.Event, concurrency)

	errors := make(chan error)
	connected := make(chan bool)

	go func() {
		for {
			c.Warning("Starting mqtt (worker: %s) on (addr: %s)...", c.WorkerName(), c.GetBrokerAddr())

			reload := make(chan bool)
			c.conn = MQTT.NewClient(opts)

			if token := c.conn.Connect(); token.Wait() && token.Error() != nil {
				errors <- fmt.Errorf(
					"Failed to establish connection with mqtt server (error: %s)",
					token.Error(),
				)
				continue
			}

			// In case we hit connected, ensure that main func is notified
			if !c.conn.IsConnected() {
				continue
			}

			c.Subscribe(c.GetBrokerTopicName(), MaxTopicSubscribeAttempts)

			// Notify rest of the app that we're ready ...
			close(connected)

			go func() {
				cct := time.Tick(2 * time.Second)

				for {
					select {
					case <-cct:
						if !c.conn.IsConnected() {
							reload <- true
							return
						}
					case <-done:
						c.Warning("Received stop signal for mqtt (worker: %s). Will not attempt to restart worker ...", c.WorkerName())
						return
					}
				}
			}()

		reloadloop:
			for {
				select {
				case <-reload:
					c.Warning("Mqtt (worker: %s) seems not to be connected. Restarting loop in 2 seconds ...", c.WorkerName())
					time.Sleep(2 * time.Second)
					break reloadloop
				}
			}

		}
	}()

	select {
	case <-connected:
		c.Info(
			"Successfully established mqtt connection for (worker: %s) on (addr: %s)",
			c.WorkerName(), c.GetBrokerAddr(),
		)
		break
	case <-time.After(time.Duration(InitialConnectionTimeout) * time.Second):
		errors <- fmt.Errorf(
			"Could not establish mqtt connection for (worker: %s) on (addr: %s) due to initial connection (timeout: %ds)",
			c.WorkerName(), c.GetBrokerAddr(), InitialConnectionTimeout,
		)
		break
	}

	return errors
}

// DrainEvents - Will return event chan back for future processing by workers
func (c *Connection) DrainEvents() chan events.Event {
	return c.events
}

// Subscribe -
func (c *Connection) Subscribe(topic string, maxRetryAttempts int) error {
	var err error

	for i := 0; i <= maxRetryAttempts; i++ {
		c.Debug(
			"About to attempt subscribe to mqtt (topic: %s) for (worker: %s) -> (retry_attempt: %d)",
			topic, c.WorkerName(), i,
		)

		if token := c.conn.Subscribe(topic, 0, nil); token.Wait() && token.Error() != nil {
			c.Error("Could not subscribe to (topic: %s) for (worker: %s) due to (err: %s). Retrying ...")
			err = token.Error()
			continue
		}

		c.Info("Successfully subscribed (worker: %s) on (topic: %s)!",
			c.WorkerName(), c.GetBrokerTopicName(),
		)

		err = nil
		break
	}

	return err
}

// BrokerHandler -
func (c *Connection) BrokerHandler(client *MQTT.Client, msg MQTT.Message) {
	c.Debug(
		"Received new mqtt (worker: %s) - (message: %s) for (topic: %s). Building event now ...",
		c.WorkerName(), msg.Payload(), msg.Topic(),
	)

	event, err := events.NewEvent(msg)

	if err != nil {
		c.Error("Could not handle received event due to (err: %s)", err)
		return
	}

	c.Debug("Event successfully created (data: %v)", event)
	c.events <- event
}

// Validate -
func (c *Connection) Validate() error {
	c.Debug("Validating mqtt configuration for (worker: %q)", c.WorkerName())

	if c.Config.Get("connection") == nil {
		return fmt.Errorf(
			"Could not validate mqtt worker as connection interface is missing (entry: %s)",
			c.Config.Get("connection"),
		)
	}

	data := c.Config.Get("connection").(map[string]interface{})

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

	if _, ok := data["username"].(string); !ok {
		return fmt.Errorf(
			"Could not validate mqtt worker as connection username is not set. Username can be empty but it MUST be set. (connection_data: %q)",
			data,
		)
	}

	if _, ok := data["password"].(string); !ok {
		return fmt.Errorf(
			"Could not validate mqtt worker as connection password is not set. Password can be empty but it MUST be set. (connection_data: %q)",
			data,
		)
	}

	clientID := data["clientId"].(string)

	if _, ok := data["clientId"].(string); !ok {
		return fmt.Errorf(
			"Could not validate mqtt worker as connection clientId is not set. (client_id: %s)",
			clientID,
		)
	}

	if len(clientID) < 2 {
		return fmt.Errorf(
			"Could not validate mqtt worker as connection clientId is not long enough. (client_id: %s)",
			clientID,
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

// GetBrokerAddr - will return full broker uri string (protocol://addr:port?params)
func (c *Connection) GetBrokerAddr() string {
	connection := c.Config.Get("connection").(map[string]interface{})
	return fmt.Sprintf("%s://%s?timeout=10s", connection["network"].(string), connection["address"].(string))
}

// GetBrokerCredentials - will return username and password defined by config
func (c *Connection) GetBrokerCredentials() (string, string) {
	connection := c.Config.Get("connection").(map[string]interface{})
	return connection["username"].(string), connection["password"].(string)
}

// GetBrokerClientId -
func (c *Connection) GetBrokerClientId() string {
	connection := c.Config.Get("connection").(map[string]interface{})
	return connection["clientId"].(string)
}

// GetBrokerTopicName -
func (c *Connection) GetBrokerTopicName() string {
	connection := c.Config.Get("connection").(map[string]interface{})
	return connection["topic"].(string)
}

// WorkerName -
func (c *Connection) WorkerName() string {
	return c.Config.Get("worker_name").(string)
}

// Stop - Will ensure that connection including subscription is killed allowing graceful timeout
func (c *Connection) Stop() error {
	c.Warning("Stopping mqtt (worker: %s) ...", c.WorkerName())

	if !c.conn.IsConnected() {
		c.Warning("Connection for mqtt (worker: %s) is already closed.", c.WorkerName())
		return nil
	}

	c.Warning("Unsubscribing from mqtt (worker: %s) (topic: %s)...", c.WorkerName(), c.GetBrokerTopicName())
	if token := c.conn.Unsubscribe(c.GetBrokerTopicName()); token.Wait() && token.Error() != nil {
		c.Error(
			"Could not unsubscribe from (topic: %s) for (worker: %s) due to (err: %s)",
			c.GetBrokerTopicName(), c.WorkerName(), token.Error(),
		)
	}

	c.Warning(
		"Stopping mqtt (worker: %s) connection (graceful_timeout: %ds)...",
		c.WorkerName(), GracefulShutdownTimeout,
	)

	c.conn.Disconnect(uint(GracefulShutdownTimeout))
	time.Sleep(time.Duration(GracefulShutdownTimeout) * time.Second)

	return nil
}
