// Copyright 2015 The PowerUnit Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package managers ...
package managers

import (
	"fmt"

	"github.com/powerunit-io/platform/logging"
)

// BaseManager -
type BaseManager struct {
	*logging.Logger

	Services map[string]Service
}

// Attach - Assing service to manager instance. Return error if service is
// already there
func (m *BaseManager) Attach(s string, i Service) error {

	if m.Exists(s) {
		return fmt.Errorf(
			"Could not attach (service: %s) as one is already attached!",
			s)
	}

	m.Services[s] = i

	return nil
}

// Remove - Remove service by its name or return error in case that service cannot
// be found within manager instance
func (m *BaseManager) Remove(s string) error {
	if !m.Exists(s) {
		return fmt.Errorf(
			"Could not remove (service: %s) as one is not assigned to the manager.",
			s)
	}

	delete(m.Services, s)
	return nil
}

// All - Return all available/attached services within manager instance
func (m *BaseManager) All() map[string]Service {
	return m.Services
}

// List - Return names of all available/attached services within manager instance
func (m *BaseManager) List() []string {
	services := []string{}

	for service := range m.Services {
		services = append(services, service)
	}

	return services
}

// Get - Return attached service or return error if it does not exist.
func (m *BaseManager) Get(s string) (Service, error) {
	if !m.Exists(s) {
		return nil, fmt.Errorf(
			"Could not retreive (service: %s) as one is not assigned to the manager.",
			s)
	}

	return m.Services[s], nil
}

// Exists - Will return boolean value based on if specific service exists
// within manager instance or not.
func (m *BaseManager) Exists(s string) bool {
	for service := range m.Services {
		if service == s {
			return true
		}
	}

	return false
}
