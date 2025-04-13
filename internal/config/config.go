package config

type Config struct {
	KUBECONFIG   string
	DEBUG        bool
	HOST         string
	API_PORT     string
	METRICS_PORT string
}

func DefaultConfig() *Config {
	return &Config{
		KUBECONFIG: GetEnvWithDefault("KUBECONFIG", "~/.kube/config"),
		DEBUG:      GetEnvBoolWithDefault("DEBUG", true),
		HOST:       GetEnvWithDefault("HOST", "0.0.0.0"),
		API_PORT: GetMultiEnvWithDefault([]string{
			"API_PORT", "PORT",
		},
			"8080",
		),
		METRICS_PORT: GetEnvWithDefault("METRICS_PORT", "8081"),
	}
}

var Instance *Config = DefaultConfig()
