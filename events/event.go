package events

import (
	MQTT "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
)

// Event -
type Event struct{}

// NewEvent -
func NewEvent(client *MQTT.Client, msg MQTT.Message) Event {
	var e Event

	return e
}
