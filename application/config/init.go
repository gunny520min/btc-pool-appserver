package config

// Load ...
func Load(configDir string) {
	LoadAppConfig(configDir + "app.yaml")
}
