package config

import (
	"os"
	"strconv"
)

func GetMultiEnvWithDefault(keys []string, defaultValue string) string {
	for _, key := range keys {
		val := os.Getenv(key)
		if val != "" {
			return val
		}
	}

	return defaultValue
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
