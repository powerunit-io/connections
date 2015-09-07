package service

import (
	"sync"

	"github.com/powerunit-io/platform/devices/device"
	"github.com/powerunit-io/platform/workers/worker"
)

// StartBinds -
func (bs *BaseService) StartBinds(wg sync.WaitGroup) {
	bs.Info("Available (binds: %v). Starting them up now ...", bs.ListAvailableBinds())

	for _, bind := range bs.GetBinds() {
		wg.Add(1)

		go func(bm BindManager) {

			if err := bm.Start(bs.done); err != nil {
				bs.Error("Could not start (bind: %s) due to (error: %s)", bm.Name(), err)
			}

			wg.Done()
		}(bind)
	}

	wg.Wait()
}

// StartDevices -
func (bs *BaseService) StartDevices(wg sync.WaitGroup) {
	bs.Info("Available (devices: %v). Starting them up now ...", bs.DevicesManager.ListAvailableDevices())

	for _, d := range bs.DevicesManager.GetDevices() {
		wg.Add(1)

		go func(d device.Device) {

			if err := d.Start(bs.done); err != nil {
				bs.Error("Could not start (device: %s) due to (error: %s)", d.DeviceName(), err)
			}
			// else {
			//	d.Handle(bs.done)
			//}

			wg.Done()
		}(d)
	}

	wg.Wait()
}

// StartWorkers -
func (bs *BaseService) StartWorkers(wg sync.WaitGroup) {
	bs.Info("Available (workers: %v). Starting them up now ...", bs.Manager.ListAvailableWorkers())

	for _, mworker := range bs.Manager.GetWorkers() {
		wg.Add(1)

		go func(w worker.Worker) {

			if err := w.Start(bs.done); err != nil {
				bs.Error("Could not start (worker: %s) due to (error: %s)", w.WorkerName(), err)
			} else {
				w.Handle(bs.done)
			}

			wg.Done()
		}(mworker)
	}

	wg.Wait()
}
