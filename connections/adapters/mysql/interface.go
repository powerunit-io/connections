// Copyright 2015 The PowerUnit Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package mysql ...
package mysql

import (
	"github.com/powerunit-io/platform/config"
	"github.com/powerunit-io/platform/logging"
	"github.com/powerunit-io/platform/managers"
)

// Adapter -
type Adapter interface {
	managers.Service
}

// NewAdapter -
func NewAdapter(n string, conf map[string]interface{}, logger *logging.Logger) (Adapter, error) {

	cnf, err := config.NewConfigManager(n, conf)

	if err != nil {
		logger.Error("Failed to configure mysql configuration manager for (manager: %s) (error: %s)", n, err)
		return nil, err
	}

	cnf.Set("name", n)

	return Adapter(&Connection{
		Logger: logger, Config: cnf, URI: cnf.Get("uri").(string),
	}), nil
}
