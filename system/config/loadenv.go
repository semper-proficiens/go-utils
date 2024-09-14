package config

import (
	"os"
	"strings"
)

// LoadEnvVars loads all environment variables into a map
func LoadEnvVars() map[string]string {
	envVars := make(map[string]string)
	for _, envVar := range os.Environ() {
		pair := strings.SplitN(envVar, "=", 2)
		key := pair[0]
		value := pair[1]
		envVars[key] = value
	}
	return envVars
}
