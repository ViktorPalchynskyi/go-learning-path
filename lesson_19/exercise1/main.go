package main

import (
	"fmt"
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port        int           `envconfig:"PORT" default:"8080"`
	DatabaseURL string        `envconfig:"DATABASE_URL" required:"true"`
	JWTSecret   string        `envconfig:"JWT_SECRET" required:"true"`
	Debug       bool          `envconfig:"DEBUG" default:"false"`
	Timeout     time.Duration `envconfig:"TIMEOUT" default:"30s"`
}

func LoadConfig() (*Config, error) {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}
	return &cfg, nil
}

func main() {
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatalf("config error: %v", err)
	}
	log.Printf("starting on port %d (debug=%v, timeout=%s)", cfg.Port, cfg.Debug, cfg.Timeout)
}
