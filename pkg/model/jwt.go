package model

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	SecretKey string `yaml:"secret_key"`
}

func JwtSecretKey() string {
	configFile, err := os.ReadFile("db.yaml")
	if err != nil {
		log.Fatalf("Failed to read db.yaml: %v", err)
	}

	var config Config
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatalf("Failed to unmarshal db.yaml: %v", err)
	}

	return config.SecretKey
}
