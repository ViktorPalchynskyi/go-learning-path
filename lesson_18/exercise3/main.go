package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

var logLevel = new(slog.LevelVar)

func newServiceLogger(service, version, env string) *slog.Logger {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: logLevel,
	})
	return slog.New(handler).With(
		"service", service,
		"version", version,
		"env", env,
	)
}

func handleCreateUser(w http.ResponseWriter, r *http.Request) {
	slog.Info("creating user", "email", "test@test.com")
	slog.Debug("debug: validating fields")
	fmt.Fprintln(w, "user created")
}

func handleListUsers(w http.ResponseWriter, r *http.Request) {
	slog.Info("listing users", "page", 1)
	slog.Debug("debug: querying database")
	fmt.Fprintln(w, "users: []")
}

func main() {
	logger := newServiceLogger("user-service", "1.0.0", "production")
	slog.SetDefault(logger)

	if os.Getenv("DEBUG") == "true" {
		logLevel.Set(slog.LevelDebug)
		slog.Info("debug mode enabled")
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/users", handleListUsers)
	mux.HandleFunc("/users/create", handleCreateUser)

	slog.Info("server started", "port", 8080)
	http.ListenAndServe(":8080", mux)
}
