package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port        int           `envconfig:"PORT" default:"8080"`
	DatabaseURL string        `envconfig:"DATABASE_URL" required:"true"`
	JWTSecret   string        `envconfig:"JWT_SECRET" required:"true"`
	Debug       bool          `envconfig:"DEBUG" default:"false"`
	Timeout     time.Duration `envconfig:"TIMEOUT" default:"30s"`
}

func init() {
	if os.Getenv("APP_ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			log.Println("no .env file found, using environment variables")
		}
	}
}

func LoadConfig() (*Config, error) {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}
	return &cfg, nil
}

func main() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	cfg, err := LoadConfig()
	if err != nil {
		log.Fatalf("config error: %v", err)
	}
	log.Printf("env=%s port=%d debug=%v", env, cfg.Port, cfg.Debug)
}
