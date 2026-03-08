package main

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type ServerConfig struct {
	Port    int           `envconfig:"SERVER_PORT" default:"8080"`
	Timeout time.Duration `envconfig:"SERVER_TIMEOUT" default:"30s"`
}

type DatabaseConfig struct {
	URL      string `envconfig:"DATABASE_URL" required:"true"`
	MaxConns int    `envconfig:"DATABASE_MAX_CONNS" default:"10"`
}

type JWTConfig struct {
	Secret string        `envconfig:"JWT_SECRET" required:"true"`
	Expiry time.Duration `envconfig:"JWT_EXPIRY" default:"24h"`
}

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
}

func init() {
	if os.Getenv("APP_ENV") != "production" {
		godotenv.Load()
	}
}

func LoadConfig() (*Config, error) {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}
	if err := validate(&cfg); err != nil {
		return nil, fmt.Errorf("config validation failed: %w", err)
	}
	return &cfg, nil
}

func validate(cfg *Config) error {
	if cfg.Server.Port <= 0 || cfg.Server.Port > 65535 {
		return fmt.Errorf("invalid port: %d", cfg.Server.Port)
	}
	if cfg.Database.URL == "" {
		return fmt.Errorf("database URL is required")
	}
	return nil
}

func logConfig(cfg *Config) {
	slog.Info("config loaded",
		"port", cfg.Server.Port,
		"timeout", cfg.Server.Timeout,
		"db_max_conns", cfg.Database.MaxConns,
		"jwt_expiry", cfg.JWT.Expiry,
	)
}

func main() {
	cfg, err := LoadConfig()
	if err != nil {
		slog.Error("failed to load config", "error", err)
		os.Exit(1)
	}
	logConfig(cfg)
}
