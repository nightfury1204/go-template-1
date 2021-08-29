package config

import (
	"log"
	"os"
	"strconv"
)

// Application holds application configurations
type Application struct {
	Host            string `yaml:"host"`
	Port            int    `yaml:"port"`
	GracefulTimeout int    `yaml:"graceful_timeout"`
}

// GetApp returns application config
func GetApp() *Application {
	appConfig := &Application{
		Host: os.Getenv("HOST"),
	}
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("invalid PORT number: %s", err)
	}
	gracefulTimeout, err := strconv.Atoi(os.Getenv("GRACEFUL_TIMEOUT"))
	if err != nil {
		log.Printf("invalid GRACEFUL_TIMEOUT: %s", err)
		gracefulTimeout = 120
	}

	appConfig.Port = port
	appConfig.GracefulTimeout = gracefulTimeout
	return appConfig
}
