package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/username/lesson_11/internal/handler"
	"github.com/username/lesson_11/internal/middleware"
	"github.com/username/lesson_11/internal/repository"
	"github.com/username/lesson_11/internal/service"
)

func main() {
	fmt.Println("Lesson 11")

	repo := repository.NewInMemoryRepo()
	service := service.NewTaskService(repo)
	handler := handler.NewHandler(service)

	r := chi.NewRouter()

	r.Use(chimiddleware.Logger)
	r.Use(chimiddleware.Recoverer)
	r.Use(middleware.LoggingMiddleware)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/tasks", func(r chi.Router) {
			r.Get("/", handler.ListTasks)
			r.Post("/", handler.CreateTask)
			r.Get("/{id}", handler.GetTask)
			r.Put("/{id}", handler.UpdateTask)
			r.Delete("/{id}", handler.DeleteTask)
		})
	})

	http.ListenAndServe(":8080", r)
}
