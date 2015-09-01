package worker

// Worker -
type Worker interface {
	String() string
	Validate() error
	Handle(done chan bool)

	Start(done chan bool) error
	Stop() error
}

// NewWorker -
func New(worker Worker) Worker {
	return worker
}
