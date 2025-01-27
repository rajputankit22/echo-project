package config

import (
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

var (
	Cfg config
)

/*
config struct - holds the configuration for the application
*/
type config struct {
	LogLevel string `env:"LOG_LEVEL" envDefault:"debug"`
	HTTPPort string `env:"HTTP_PORT" envDefault:":8090"`
	Cache cache
}

// Cache - holds the configuration for the cache
type cache struct {
	Host     string `env:"CACHE_HOST" envDefault:"localhost"`
	Port     string `env:"CACHE_PORT" envDefault:"6379"`
	Password string `env:"CACHE_PASSWORD" envDefault:""`
	DB       int    `env:"CACHE_DB" envDefault:"0"`
}


func Init() *config {
	// Load the .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Parse environment variables into the st ruct
	if err := env.Parse(&Cfg); err != nil {
		log.Fatalf("Failed to load environment variables: %v", err)
	}

	// Load the configuration from the environment
	return &Cfg
}

func Config() *config {
	if Cfg.LogLevel == "" {
		return Init()
	}
	return &Cfg
}
