package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

type contextKey struct{}

func setupLogger() {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
	slog.SetDefault(slog.New(handler))
}

func RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get("X-Request-ID")
		if requestID == "" {
			requestID = fmt.Sprintf("%d", time.Now().UnixNano())
		}

		logger := slog.Default().With("request_id", requestID)
		ctx := context.WithValue(r.Context(), contextKey{}, logger)

		w.Header().Set("X-Request-ID", requestID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func LoggerFromCtx(ctx context.Context) *slog.Logger {
	if l, ok := ctx.Value(contextKey{}).(*slog.Logger); ok {
		return l
	}
	return slog.Default()
}

func handleGetUser(w http.ResponseWriter, r *http.Request) {
	log := LoggerFromCtx(r.Context())
	log.Info("getting user", "user_id", 42)
	fmt.Fprintln(w, "user: Viktor")
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	log := LoggerFromCtx(r.Context())
	log.Info("health check")
	fmt.Fprintln(w, "ok")
}

func main() {
	setupLogger()

	mux := http.NewServeMux()
	mux.HandleFunc("/user", handleGetUser)
	mux.HandleFunc("/health", handleHealth)

	slog.Info("server starting", "port", 8080)
	http.ListenAndServe(":8080", RequestIDMiddleware(mux))
}
