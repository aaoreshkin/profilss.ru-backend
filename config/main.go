package config

import (
	"log"
	"os"
	"sync"
)

type (
	Config struct {
		Host, Port, Path, SecretKey string
		Database                    Database
	}

	Database struct {
		URL string
	}
)

var (
	config *Config
	once   sync.Once
)

// Initializes a new Config instance.
func NewConfig() (*Config, error) {
	once.Do(func() {
		config = &Config{
			Host:      env("HOST"),
			Port:      env("PORT"),
			Path:      env("REMOTE_PATH"),
			SecretKey: env("SECRET_KEY"),
			Database: Database{
				URL: env("DB_URL"),
			},
		}
	})
	return config, nil
}

func env(key string) string {
	value, ok := os.LookupEnv(key)
	// If the variable was not set, return error.
	if !ok {
		log.Fatal("Could not load environment variable")
	}
	// Return the value of the variable.
	return value
}
