package service

import (
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/powerunit-io/platform/config"
	"github.com/powerunit-io/platform/logging"
	"github.com/powerunit-io/platform/utils"
)

// BaseService -
type BaseService struct {
	*logging.Logger
	Config *config.ConfigManager

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
		bs.Config.Get("ServiceName"), bs.Config.Get("ServiceVersion"),
	)

	bs.done = make(chan bool)

	go bs.HandleSigterm()

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

	bs.Warning("Service (name: %s) is now stopped!", bs.Config.Get("ServiceName"))
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
