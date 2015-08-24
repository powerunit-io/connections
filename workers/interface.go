package workers

type Worker interface{}

// NewWorker -
func NewWorker(worker Worker) *Worker {
	return &worker
}
