// Copyright 2015 The PowerUnit Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package config ...
package config

func init() {
	// Make sure that ConfigManager are actually created before we process with anything ....
	if ConfigManager == nil {
		ConfigManager = make(map[string]interface{})
	}
}
