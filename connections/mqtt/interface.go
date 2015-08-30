package mqtt

import (
	"github.com/powerunit-io/platform/config"
	"github.com/powerunit-io/platform/logging"
)

type MqttManager interface {
	Start(done chan bool) error
	Validate() error
	String() string
	Stop() error
}

func NewConnection(log *logging.Logger, config *config.Config) (MqttManager, error) {
	manager := func(m MqttManager) MqttManager {
		return m
	}

	connection := MqttConnection{Logger: log, Config: config}

	return manager(&connection), nil
}
