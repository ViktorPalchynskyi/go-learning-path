package main

import (
	"fmt"
	"log/slog"
	"os"
)

func setupLogger(level slog.Level) {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     level,
		AddSource: true,
	})
	slog.SetDefault(slog.New(handler))
}

func main() {
	setupLogger(slog.LevelInfo)

	slog.Info("server starting", "port", 8080, "env", "production")
	slog.Warn("slow query", "duration_ms", 450, "query", "SELECT * FROM users")

	err := fmt.Errorf("connection refused")
	slog.Error("database connection failed", "error", err, "host", "localhost")

	slog.Debug("this should not appear", "detail", "hidden at INFO level")
}
