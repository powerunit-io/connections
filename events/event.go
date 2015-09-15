// Copyright 2015 The PowerUnit Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package events ...
package events

import (
	"bytes"
	"encoding/json"
	"fmt"

	MQTT "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
	"github.com/powerunit-io/platform/utils"
)

// Event -
type Event struct {
	MQTT.Message
	EventType string                 `json:"type"`
	DeviceID  string                 `json:"device_id"`
	Data      map[string]interface{} `json:"data"`
}

// Validate -
func (e *Event) Validate() error {

	if !utils.StringInSlice(e.EventType, AvailableEventTypes) {
		return fmt.Errorf(
			"Could not validate (event: %v). (event_type: %s) is not under available (event_types: %v)",
			e, e.EventType, AvailableEventTypes,
		)
	}

	return nil
}

// NewEvent -
func NewEvent(msg MQTT.Message) (Event, error) {
	e := Event{Message: msg}

	decoder := json.NewDecoder(bytes.NewReader(msg.Payload()))

	if err := decoder.Decode(&e); err != nil {
		return e, err
	}

	if err := e.Validate(); err != nil {
		return e, err
	}

	return e, nil
}
