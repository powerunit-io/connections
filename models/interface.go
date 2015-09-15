// Copyright 2015 The PowerUnit Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package models ...
package models

// Model -
type Model interface {
	List() []Model
	Get() Model
	Exists() bool
	Create(data map[string]interface{}) Model
	Update(data map[string]interface{}) Model
	Delete() bool
}
