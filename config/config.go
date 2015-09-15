// Copyright 2015 The PowerUnit Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package config ...
package config

import "github.com/powerunit-io/platform/utils"

// Config - Configuration manager helper designed to address configuration items
type Config struct {
	Config map[string]interface{}
}

// Set - Will set value of requested key within configuration manager instance
func (c *Config) Set(key string, value interface{}) {
	c.Config[key] = value
}

// SetMany - Will loop through map and set string -> interface{}
func (c *Config) SetMany(config map[string]interface{}) {
	for key, value := range config {
		c.Set(key, value)
	}
}

// Get - Retreive configuration manager config value by key
func (c *Config) Get(key string) interface{} {
	if !c.KeyExists(key) {
		return nil
	}

	return c.Config[key]
}

// KeyExists - Check whenever key exists within configuration manager instance
func (c *Config) KeyExists(key string) bool {
	return utils.KeyInSlice(key, c.Config)
}
