package worker

import "github.com/powerunit-io/platform/events"

// Worker -
type Worker interface {
	String() string
	Validate() error
	Handle(e *events.Event) error

	Start(done chan bool) error
	Stop() error
}

// NewWorker -
func NewWorker(worker Worker) Worker {
	return worker
}
