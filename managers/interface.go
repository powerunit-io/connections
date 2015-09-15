// Copyright 2015 The PowerUnit Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package managers ...
package managers

// Service -
type Service interface {
	Start(done chan bool) error
	Stop() error
	Validate() error
	Name() string
}

// Manager -
type Manager interface {
	Attach(m string, bm Service) error
	Remove(m string) error
	All() map[string]Service
	List() []string
	Get(m string) (Service, error)
	Exists(m string) bool
}
