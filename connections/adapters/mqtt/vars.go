// Copyright 2015 The PowerUnit Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package mqtt ...
package mqtt

var (
	// AvailableConnectionTypes -
	AvailableConnectionTypes = []string{"tcp", "tls", "ws"}

	// InitialConnectionTimeout -
	InitialConnectionTimeout = 10

	// MaxTopicSubscribeAttempts -
	MaxTopicSubscribeAttempts = 5

	// GracefulShutdownTimeout -
	GracefulShutdownTimeout = 1
)
