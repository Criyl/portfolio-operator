package config

import (
        "os"
        "testing"
        "github.com/stretchr/testify/assert"
)

func TestGetMultiEnvWithDefault(t *testing.T) {
        os.Setenv("KEY1", "value1")
        os.Setenv("KEY2", "")

        result := GetMultiEnvWithDefault([]string{"KEY1", "KEY2"}, "default")
        assert.Equal(t, "value1", result)

        result = GetMultiEnvWithDefault([]string{"KEY3", "KEY4"}, "default")
        assert.Equal(t, "default", result)
}

func TestGetEnvWithDefault(t *testing.T) {
        os.Setenv("KEY1", "value1")

        result := GetEnvWithDefault("KEY1", "default")
        assert.Equal(t, "value1", result)

        os.Unsetenv("KEY1")

        result = GetEnvWithDefault("KEY1", "default")
        assert.Equal(t, "default", result)
}

func TestGetEnvBoolWithDefault(t *testing.T) {
        os.Setenv("KEY1", "true")

        result := GetEnvBoolWithDefault("KEY1", false)
        assert.True(t, result)

        os.Unsetenv("KEY1")

        result = GetEnvBoolWithDefault("KEY1", false)
        assert.False(t, result)
}