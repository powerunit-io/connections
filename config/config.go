package config

import "github.com/powerunit-io/platform/utils"

// Config -
type Config struct {
	Config map[string]interface{}
}

// Set -
func (c *Config) Set(key string, value interface{}) {
	c.Config[key] = value
}

// SetMany -
func (c *Config) SetMany(config map[string]interface{}) {
	for key, value := range config {
		c.Set(key, value)
	}
}

// Get -
func (c *Config) Get(key string) interface{} {
	if !c.KeyExists(key) {
		return nil
	}

	return c.Config[key]
}

// KeyExists -
func (c *Config) KeyExists(key string) bool {
	return utils.KeyInSlice(key, c.Config)
}
