package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/username/lesson_11/internal/domain"
	"github.com/username/lesson_11/internal/service"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type CreateTaskRequest struct {
	Title string `json:"title"`
}

type UpdateTaskRequest struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type Handler struct {
	service *service.TaskService
}

func NewHandler(service *service.TaskService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var req CreateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, 400, ErrorResponse{Error: "invalid JSON"})
		return
	}
	task, err := h.service.CreateTask(req.Title)
	if err != nil {
		writeJSON(w, 500, ErrorResponse{Error: "internal error"})
		return
	}
	writeJSON(w, 201, task)
}

func (h *Handler) ListTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.service.ListTasks()
	if err != nil {
		writeJSON(w, 500, ErrorResponse{Error: "internal error"})
		return
	}
	writeJSON(w, 200, tasks)
}

func (h *Handler) GetTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	task, err := h.service.GetTask(id)
	if err != nil {
		if isNotFound(err) {
			writeJSON(w, 404, ErrorResponse{Error: "task not found"})
			return
		}
		writeJSON(w, 500, ErrorResponse{Error: "internal error"})
		return
	}
	writeJSON(w, 200, task)
}

func (h *Handler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var req UpdateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, 400, ErrorResponse{Error: "invalid JSON"})
		return
	}
	task, err := h.service.UpdateTask(id, req.Title, req.Completed)
	if err != nil {
		if isNotFound(err) {
			writeJSON(w, 404, ErrorResponse{Error: "task not found"})
			return
		}
		writeJSON(w, 500, ErrorResponse{Error: "internal error"})
		return
	}
	writeJSON(w, 200, task)
}

func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := h.service.DeleteTask(id); err != nil {
		if isNotFound(err) {
			writeJSON(w, 404, ErrorResponse{Error: "task not found"})
			return
		}
		writeJSON(w, 500, ErrorResponse{Error: "internal error"})
		return
	}
	w.WriteHeader(204)
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func isNotFound(err error) bool {
	return errors.Is(err, domain.ErrNotFound)
}
