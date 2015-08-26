package manager

import (
	"fmt"

	"github.com/powerunit-io/platform/logging"
	"github.com/powerunit-io/platform/workers/worker"
)

// WorkerManager -
type WorkerManager struct {
	*logging.Logger
	Workers map[string]worker.Worker
}

// GetWorkers
func (wm *WorkerManager) GetWorkers() map[string]worker.Worker {
	return wm.Workers
}

// AttachWorker -
func (wm *WorkerManager) AttachWorker(wn string, w worker.Worker) error {

	if wm.WorkerExists(wn) {
		return fmt.Errorf("Could not attach (worker: %s) as one is already attached!", wn)
	}

	// Make sure that all is validated full before we go anywhere with it ...
	if err := w.Validate(); err != nil {
		return err
	}

	wm.Workers[wn] = w

	return nil
}

// RemoveWorker -
func (wm *WorkerManager) RemoveWorker(wn string) error {
	if !wm.WorkerExists(wn) {
		return fmt.Errorf("Could not remove (worker: %s) as one does not exist in store...", wn)
	}

	delete(wm.Workers, wn)
	return nil
}

// ListAvailableWorkers -
func (wm *WorkerManager) ListAvailableWorkers() []string {
	workers := []string{}

	for worker, _ := range wm.Workers {
		workers = append(workers, worker)
	}

	return workers
}

// GetWorker -
func (wm *WorkerManager) GetWorker(wn string) (worker.Worker, error) {
	if !wm.WorkerExists(wn) {
		return nil, fmt.Errorf("Could not retreive (worker: %s) as one does not exist in store...", wn)
	}

	return wm.Workers[wn], nil
}

// WorkerExists -
func (wm *WorkerManager) WorkerExists(wn string) bool {
	for worker, _ := range wm.Workers {
		if worker == wn {
			return true
		}
	}

	return false
}
