package config

import "github.com/powerunit-io/platform/utils"

// ConfigManagers -
var ConfigManagers map[string]interface{}

// GetConfigManager -
func GetConfigManager(managerName string) *ConfigManager {
	manager := ConfigManagers[managerName].(ConfigManager)
	return &manager
}

// -----------------------------------------------------------------------------

// ConfigManager -
type ConfigManager struct {
	Config map[string]interface{}
}

// Set -
func (cm *ConfigManager) Set(key string, value interface{}) {
	cm.Config[key] = value
}

// SetMany -
func (cm *ConfigManager) SetMany(config map[string]interface{}) {
	for key, value := range config {
		cm.Set(key, value)
	}
}

// Get -
func (cm *ConfigManager) Get(key string) interface{} {
	if !cm.KeyExists(key) {
		return nil
	}

	return cm.Config[key]
}

// KeyExists -
func (cm *ConfigManager) KeyExists(key string) bool {
	return utils.KeyInSlice(key, cm.Config)
}

// NewConfigManager -
func NewConfigManager(managerName string, managerConfig map[string]interface{}) *ConfigManager {
	if !utils.KeyInSlice(managerName, ConfigManagers) {
		ConfigManagers[managerName] = ConfigManager{
			Config: managerConfig,
		}
	}

	return GetConfigManager(managerName)
}

func init() {
	// Make sure that ConfigManagers are actually created before we process with anything ....
	if ConfigManagers == nil {
		ConfigManagers = make(map[string]interface{})
	}
}
