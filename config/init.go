package config

func init() {
	// Make sure that ConfigManager are actually created before we process with anything ....
	if ConfigManager == nil {
		ConfigManager = make(map[string]interface{})
	}
}
