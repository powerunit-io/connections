// Copyright 2015 The PowerUnit Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package workers ...
package workers

import (
	"github.com/powerunit-io/platform/config"
	"github.com/powerunit-io/platform/logging"
)

// WorkerBase -
type WorkerBase struct {
	*logging.Logger
	*config.Config
}

// Name - Will return name of the worker ...
func (wb *WorkerBase) Name() string {
	return wb.Config.Get("name").(string)
}
