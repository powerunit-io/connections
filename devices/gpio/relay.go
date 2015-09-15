// Copyright 2015 The PowerUnit Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package gpio ...
package gpio

import "github.com/powerunit-io/platform/config"

// Relay -
type Relay struct {
	*config.Config
}

// Pin -
func (s *Relay) Pin() uint64 {
	return 0
}

// Name - Will return name of this relay
func (s *Relay) Name() string {
	return s.Config.Get("name").(string)
}
