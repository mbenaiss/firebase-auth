package config

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// Config is the configuration of the application.
type Config struct {
	Port                         string `envconfig:"PORT"`
	FirebaseAPIKey               string `envconfig:"FIREBASE_API_KEY"`
	GoogleApplicationCredentials string `envconfig:"GOOGLE_APPLICATION_CREDENTIALS"`
}

// New retrun new instance of Config.
func New() (*Config, error) {
	c := Config{
		Port: "8080",
	}

	err := godotenv.Load()
	if err != nil {
		log.Println("unable to load .env file")
	}

	err = envconfig.Process("", &c)
	if err != nil {
		return nil, fmt.Errorf("unable to get envconfig %w", err)
	}

	return &c, nil
}
