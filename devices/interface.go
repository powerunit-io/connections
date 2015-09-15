// Copyright 2015 The PowerUnit Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package devices ...
package devices

import (
	"github.com/powerunit-io/platform/logging"
	"github.com/powerunit-io/platform/managers"
)

// Device -
type Device interface {
	Name() string
	Pin() uint64
}

// Manager -
type Manager interface {
	managers.Manager
}

// NewManager -
func NewManager(logger *logging.Logger) Manager {
	return Manager(&managers.BaseManager{Logger: logger})
}
