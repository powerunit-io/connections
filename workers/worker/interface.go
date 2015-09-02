package worker

import "github.com/powerunit-io/platform/events"

// Worker -
type Worker interface {
	Start(done chan bool) error
	Stop() error

	Validate() error

	Drain() chan events.Event
	Handle(done chan bool)

	WorkerName() string
}

// New -
func New(worker Worker) Worker {
	return worker
}
