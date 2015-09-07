package service

import (
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"

	"github.com/powerunit-io/platform/config"
	"github.com/powerunit-io/platform/devices/device"
	devices "github.com/powerunit-io/platform/devices/manager"
	helpers "github.com/powerunit-io/platform/helpers/manager"
	"github.com/powerunit-io/platform/logging"
	"github.com/powerunit-io/platform/utils"
	"github.com/powerunit-io/platform/workers/manager"
	"github.com/powerunit-io/platform/workers/worker"
	"github.com/telapi/flumine/logger"
)

// BaseService -
type BaseService struct {
	*logging.Logger
	Config         *config.Config
	Manager        manager.Manager
	HelpersManager helpers.Manager
	DevicesManager devices.Manager
	Binds          map[string]BindManager
	done           chan bool
}

// SetGoMaxProcs - Will set GOMAXPROCS based on custom ENV key or NumCPU
func (bs *BaseService) SetGoMaxProcs(envKey string) {
	numCPU := utils.GetProcessCount(envKey)
	bs.Warning(
		"Setting up runtime in parallel mode -> (threads: %d) ...", numCPU,
	)
	runtime.GOMAXPROCS(numCPU)
}

// Start - Generic service start function
func (bs *BaseService) Start() error {
	bs.Info(
		"Starting up (service: %s) - (ver: %v)",
		bs.Config.Get("service_name"), bs.Config.Get("service_version"),
	)

	bs.done = make(chan bool)

	var wg sync.WaitGroup

	go bs.HandleSigterm()

	bs.StartBinds(wg)

	// Will basically import devices from devices manager...
	if err := bs.DevicesManager.ImportDevices(); err != nil {
		logger.Error("Could not import devices due to (error: %s)", err)
		close(bs.done)
	} else {
		bs.StartDevices(wg)
		bs.StartWorkers(wg)
	}

	select {
	case <-bs.done:
		bs.Warning("Service exit signal caught. Waiting for 5 seconds and killing...")

		// Give it a little bit of the time.
		time.Sleep(0 * time.Second)

		bs.Stop()
	}

	// I know ... BUT i want to be able return error on start so deal with it :)
	return nil
}

// Stop - Generic service stop function
func (bs *BaseService) Stop() error {
	var wg sync.WaitGroup

	for _, mbind := range bs.GetBinds() {
		wg.Add(1)

		go func(bm BindManager) {
			if err := bm.Stop(); err != nil {
				bs.Error("Could not stop (bind: %s) due to (error: %s)", bm.Name(), err)
			}
			wg.Done()
		}(mbind)
	}

	wg.Wait()

	for _, mworker := range bs.Manager.GetWorkers() {
		wg.Add(1)

		go func(w worker.Worker) {
			if err := w.Stop(); err != nil {
				bs.Error("Could not stop (worker: %s) due to (error: %s)", w.WorkerName(), err)
			}
			wg.Done()
		}(mworker)
	}

	wg.Wait()

	for _, mdevice := range bs.DevicesManager.GetDevices() {
		wg.Add(1)

		go func(d device.Device) {
			if err := d.Stop(); err != nil {
				bs.Error("Could not stop (device: %s) due to (error: %s)", d.DeviceName(), err)
			}
			wg.Done()
		}(mdevice)
	}

	wg.Wait()

	bs.Warning("Service (name: %s) is now stopped!", bs.Config.Get("service_name"))
	os.Exit(0)

	return nil
}

// HandleSigterm - Will basically wait for channel to close and than initiate
// service stop logic followed by actual exit.
func (bs *BaseService) HandleSigterm() {
	skill := make(chan os.Signal, 1)

	signal.Notify(skill, os.Interrupt)
	signal.Notify(skill, syscall.SIGTERM)

	<-skill

	close(bs.done)
}
