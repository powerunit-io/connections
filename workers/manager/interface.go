package manager

import (
	"github.com/powerunit-io/platform/logging"
	"github.com/powerunit-io/platform/workers/worker"
)

// Manager -
type Manager interface {
	GetWorkers() map[string]worker.Worker
	AttachWorker(wn string, w worker.Worker) error
	RemoveWorker(wn string) error
	ListAvailableWorkers() []string
	GetWorker(wn string) (worker.Worker, error)
	WorkerExists(wn string) bool
}

// NewWorkerManager -
func NewWorkerManager(log *logging.Logger) Manager {
	return Manager(&WorkerManager{
		Logger:  log,
		Workers: make(map[string]worker.Worker),
	})
}
