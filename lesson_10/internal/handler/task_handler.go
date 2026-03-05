package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/username/lesson_10/internal/repository"
	"github.com/username/lesson_10/internal/service"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type CreateTaskRequest struct {
	Title string
}

type Handler struct {
	service *service.TaskService
}

func NewHandler(service *service.TaskService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	task, err := h.service.GetTask(ctx, id)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			writeJSON(w, 504, ErrorResponse{Error: "request timeout"})
			return
		}
		if errors.Is(err, repository.ErrNotFound) {
			writeJSON(w, 404, ErrorResponse{Error: "task not found"})
			return
		}
		writeJSON(w, 500, ErrorResponse{Error: "internal error"})
		return
	}

	writeJSON(w, 200, task)
}

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var req CreateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, 400, ErrorResponse{Error: "invalid JSON"})
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	task, err := h.service.CreateTask(ctx, req.Title)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			writeJSON(w, 504, ErrorResponse{Error: "request timeout"})
			return
		}
		writeJSON(w, 500, ErrorResponse{Error: "internal error"})
		return
	}

	writeJSON(w, 201, task)
}

func writeJSON(w http.ResponseWriter, status int, data any)  {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
