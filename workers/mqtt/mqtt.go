package mqtt

import (
	"fmt"
	"time"

	"github.com/powerunit-io/platform/config"
	"github.com/powerunit-io/platform/connections/mqtt"
	"github.com/powerunit-io/platform/events"
	"github.com/powerunit-io/platform/logging"
	"github.com/powerunit-io/platform/workers/worker"
)

// Worker -
type Worker struct {
	*logging.Logger
	MqttConnection mqtt.Connection
	worker.BaseWorker
}

// Start -
func (w *Worker) Start(done chan bool) error {
	w.Warning("Starting up (worker: %s) ...", w.String())

	errors := w.MqttConnection.Start(done)

	// Just one error for now ...
	select {
	case err := <-errors:
		return fmt.Errorf(
			"Failed to start mqtt connection for (worker: %s) due to (err: %s)",
			w.String(), err,
		)
	case <-time.After(2 * time.Second):
		break
	}

	return nil
}

// Validate -
func (w *Worker) Validate() error {

	// Validate basic stuff ...
	if err := w.BaseWorker.Validate(); err != nil {
		return err
	}

	// Validate MQTT connection requirements ...
	if err := w.MqttConnection.Validate(); err != nil {
		return err
	}

	return nil
}

// Stop -
func (w *Worker) Stop() error {
	w.Warning("Stopping (worker: %s) ...", w.String())

	// Stopping MqttConnection worker ...
	w.MqttConnection.Stop()

	return nil
}

// Drain - Just proxy to MqttConnection DrainEvents()
func (w *Worker) Drain() chan events.Event {
	return w.MqttConnection.DrainEvents()
}

// NewWorker - Will create new MQTT related worker. Error will be return in case of any issues
func NewWorker(name string, log *logging.Logger, conf *config.Config) (*Worker, error) {
	return &Worker{
		Logger:         log,
		BaseWorker:     worker.BaseWorker{Logger: log, Config: conf},
		MqttConnection: mqtt.Connection{Logger: log, Config: conf},
	}, nil
}
