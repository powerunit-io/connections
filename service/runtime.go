package service

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/powerunit-io/platform/managers"
)

// Start -
func (bs *BaseService) Start() error {
	bs.Info("Starting up (service: %s) - (ver: %v)", bs.Name(), bs.Config.Get("service_version"))

	go bs.HandleSigterm()

	if err := bs.StartDevices(); err != nil {
		return err
	}

	if err := bs.StartWorkers(); err != nil {
		return err
	}

	select {
	case <-bs.Done:
		bs.Warning("Service exit signal caught. Waiting for 5 seconds and killing...")

		// Give it a little bit of the time.
		time.Sleep(0 * time.Second)

		bs.Stop()
	}

	// I know ... BUT i want to be able return error on start so deal with it :)
	return nil
}

// Stop -
func (bs *BaseService) Stop() error {

	for _, service := range bs.Connections.All() {
		go func(s managers.Service) {
			if err := s.Stop(); err != nil {
				bs.Error("Could not stop (connection: %s) due to (error: %s)", s.Name(), err)
			}
		}(service)
	}

	for _, service := range bs.Devices.All() {
		go func(s managers.Service) {
			if err := s.Stop(); err != nil {
				bs.Error("Could not stop (device: %s) due to (error: %s)", s.Name(), err)
			}
		}(service)
	}

	for _, service := range bs.Workers.All() {
		go func(s managers.Service) {
			if err := s.Stop(); err != nil {
				bs.Error("Could not stop (worker: %s) due to (error: %s)", s.Name(), err)
			}
		}(service)
	}

	bs.Warning("Service (name: %s) is now stopped!", bs.Name())
	os.Exit(0)

	return nil
}

// StartConnections -
func (bs *BaseService) StartConnections() (err error) {
	var wg sync.WaitGroup

	bs.Info("Available (connections: %v). Starting them up now ...", bs.Connections.List())

	for _, service := range bs.Connections.All() {
		wg.Add(1)

		go func(s managers.Service) {

			if err = s.Start(bs.Done); err != nil {
				bs.Error("Could not start (connection: %s) due to (error: %s)", s.Name(), err)
			}

			wg.Done()
		}(service)
	}

	wg.Wait()

	return
}

// StartDevices -
func (bs *BaseService) StartDevices() (err error) {
	var wg sync.WaitGroup

	bs.Info("Available (devices: %v). Starting them up now ...", bs.Devices.List())

	for _, service := range bs.Devices.All() {
		wg.Add(1)

		go func(s managers.Service) {

			if err = s.Start(bs.Done); err != nil {
				bs.Error("Could not start (device: %s) due to (error: %s)", s.Name(), err)
			}

			wg.Done()
		}(service)
	}

	wg.Wait()
	return
}

// StartWorkers -
func (bs *BaseService) StartWorkers() (err error) {
	var wg sync.WaitGroup

	bs.Info("Available (workers: %v). Starting them up now ...", bs.Workers.List())

	for _, service := range bs.Workers.All() {
		wg.Add(1)

		go func(s managers.Service) {

			if err = s.Start(bs.Done); err != nil {
				bs.Error("Could not start (worker: %s) due to (error: %s)", s.Name(), err)
			}

			wg.Done()
		}(service)
	}

	wg.Wait()
	return
}

// HandleSigterm - Will basically wait for channel to close and than initiate
// service stop logic followed by actual exit.
func (bs *BaseService) HandleSigterm() {
	skill := make(chan os.Signal, 1)

	signal.Notify(skill, os.Interrupt)
	signal.Notify(skill, syscall.SIGTERM)

	<-skill

	close(bs.Done)
}
