// Copyright 2015 The PowerUnit Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package models ...
package models

import "time"

// BaseModel -
type BaseModel struct {
	ID        int64 `sql:"AUTO_INCREMENT",gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time `sql:"DEFAULT:current_timestamp"`
	DeletedAt *time.Time
}
