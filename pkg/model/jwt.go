package model

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	SecretKey string `yaml:"JWTSecretKey"`
}

func JwtSecretKey() string {
	configFile, err := os.ReadFile("db.yaml")
	if err != nil {
		fmt.Println("1")
		log.Fatalf("Failed to read db.yaml: %v", err)
	}

	var config Config
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		fmt.Println("2")
		log.Fatalf("Failed to unmarshal db.yaml: %v", err)
	}

	return config.SecretKey
}
