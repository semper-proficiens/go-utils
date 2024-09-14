package config

import (
	"log"
	"testing"
)

func TestLoadEnvVars(t *testing.T) {
	t.Setenv("1_KEY", "1_value")
	t.Setenv("2_KEY", "2_value")
	myEnvVars := LoadEnvVars()
	for k, v := range myEnvVars {
		if k == "1_KEY" || k == "2_KEY" {
			log.Println("key:", k, "value:", v)
		}
	}
}
