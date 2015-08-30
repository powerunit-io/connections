package service

import (
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"

	"github.com/powerunit-io/platform/config"
	"github.com/powerunit-io/platform/logging"
	"github.com/powerunit-io/platform/utils"
	"github.com/powerunit-io/platform/workers/manager"
	"github.com/powerunit-io/platform/workers/worker"
)

// BaseService -
type BaseService struct {
	*logging.Logger
	*config.Config
	manager.Manager
	done chan bool
}

// SetGoMaxProcs - Will set GOMAXPROCS based on custom ENV key or NumCPU
func (bs *BaseService) SetGoMaxProcs(envKey string) {
	numCpu := utils.GetProcessCount(envKey)
	bs.Warning(
		"Setting up runtime in parallel mode -> (threads: %d) ...", numCpu,
	)
	runtime.GOMAXPROCS(numCpu)
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

	bs.Info("Available (workers: %v). Starting them up now ...", bs.Manager.ListAvailableWorkers())

	for _, mworker := range bs.Manager.GetWorkers() {
		wg.Add(1)

		go func(w worker.Worker) {
			if err := w.Start(bs.done); err != nil {
				bs.Error("Could not start (worker: %s) due to (error: %s)", w.String(), err)
			}
			wg.Done()
		}(mworker)
	}

	wg.Wait()

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

	for _, mworker := range bs.Manager.GetWorkers() {
		wg.Add(1)

		go func(w worker.Worker) {
			if err := w.Stop(); err != nil {
				bs.Error("Could not stop (worker: %s) due to (error: %s)", w.String(), err)
			}
			wg.Done()
		}(mworker)
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
