package config

import (
	"log"

	"gopkg.in/parse.v2"
	yaml "gopkg.in/yaml.v2"
)

// Config .
type Config struct {
	Port      string `yaml:"port"`
	Driver    string `yaml:"driver"`
	DBConnStr string `yaml:"db"`
}

// Parse .
func Parse(path string) *Config {
	config := Config{}
	if err := parse.File(&config, path, yaml.Unmarshal); err != nil {
		log.Fatal(err)
	}

	return &config
}
