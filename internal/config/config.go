package config

import (
	"os"
	"strconv"
)

type Config struct {
	HOST       string
	PORT       string
	KUBECONFIG string
	DEBUG      bool
}

func GetEnvWithDefault(key string, defaultValue string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultValue
	}
	return val
}

func GetEnvBoolWithDefault(key string, defaultValue bool) bool {
	val := os.Getenv(key)
	if val == "" {
		return defaultValue
	}

	boolVal, err := strconv.ParseBool(val)
	if err != nil {
		return defaultValue
	}
	return boolVal
}

var Instance *Config = &Config{
	KUBECONFIG: GetEnvWithDefault("KUBECONFIG", "~/.kube/config"),
	HOST:       GetEnvWithDefault("HOST", "0.0.0.0"),
	PORT:       GetEnvWithDefault("PORT", "8080"),
	DEBUG:      GetEnvBoolWithDefault("DEBUG", true),
}
