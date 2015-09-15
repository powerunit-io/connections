// Copyright 2015 The PowerUnit Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package models ...
package models

// BaseModel -
type BaseModel struct {
	Id int64
}

// List -
func (bm *BaseModel) List() []Model {
	return nil
}

// Get -
func (bm *BaseModel) Get() Model {
	return nil
}

// Exists -
func (bm *BaseModel) Exists() bool {
	return false
}

// Create -
func (bm *BaseModel) Create(data map[string]interface{}) Model {
	return nil
}

// Update -
func (bm *BaseModel) Update(data map[string]interface{}) Model {
	return nil
}

// Delete -
func (bm *BaseModel) Delete() bool {
	return false
}
