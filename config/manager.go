package config

import (
	"fmt"

	"github.com/powerunit-io/platform/utils"
)

// ConfigManagers -
var ConfigManager map[string]interface{}

// ConfigManagerExists -
func ConfigManagerExists(manager string) bool {
	return utils.KeyInSlice(manager, ConfigManager)
}

// GetConfigManager -
func GetConfigManager(managerName string) (*Config, error) {

	if !ConfigManagerExists(managerName) {
		return nil, fmt.Errorf("Could not discover configuration (manager: %s). Forgot to load it?", managerName)
	}

	manager := ConfigManager[managerName].(Config)
	return &manager, nil
}

// SetConfigManager -
func SetConfigManager(managerName string, configData map[string]interface{}) (*Config, error) {
	if !ConfigManagerExists(managerName) {
		ConfigManager[managerName] = Config{
			Config: configData,
		}
	}

	return GetConfigManager(managerName)
}

// -----------------------------------------------------------------------------

// NewConfigManager -
func NewConfigManager(managerName string, configData map[string]interface{}) (*Config, error) {
	return SetConfigManager(managerName, configData)
}
