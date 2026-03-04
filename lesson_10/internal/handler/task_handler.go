package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/username/lesson_10/internal/service"
)

type ErrorResponse struct {
	Error string
}

type CreateTaskRequest struct {
	Title string
}

type Handler struct {
	service service.TaskService
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) GetTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	task, err := h.service.GetTask(ctx, id)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			writeJSON(w, 500, ErrorResponse{Error: "invalid JSON"})
			return
		}
	}

	writeJSON(w, 202, task)
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
			writeJSON(w, 500, ErrorResponse{Error: "invalid JSON"})
			return
		}
	}

	writeJSON(w, 202, task)
}

func writeJSON(w http.ResponseWriter, status int, data any)  {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
