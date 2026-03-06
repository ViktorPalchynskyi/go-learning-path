package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/username/lesson_13/internal/domain"
	"github.com/username/lesson_13/internal/service"
)

type CreateUserRequest struct {
	Name string
	Email string
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler{
	return &UserHandler{service: service}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, 400, ErrorResponse{Error: "invalid JSON"})
		return
	}

	user, err := h.service.Create(ctx, req.Name, req.Email)
	if err != nil {
		writeJSON(w, 500, ErrorResponse{Error: "internal error"})
		return
	}

	writeJSON(w, 201, user)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64);
	if err !=nil {
		writeJSON(w, 400, ErrorResponse{Error: "invalid id"})
		return
	}

	user, err := h.service.FindById(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			writeJSON(w, 404, ErrorResponse{Error: "not found"})
			return
		}
		writeJSON(w, 500, ErrorResponse{Error: "internal error"})
		return
	}
	writeJSON(w, 200, user)
}

func (h *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	users, err := h.service.FindAll(ctx)
	if err != nil {
		writeJSON(w, 500, ErrorResponse{Error: "internal error"})
		return
	}
	writeJSON(w, 200, users)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64);
	if err !=nil { 
		writeJSON(w, 400, ErrorResponse{Error: "invalid id"})
		return
	}
	if err := h.service.Delete(ctx, id); err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			writeJSON(w, 404, ErrorResponse{Error: "not found"})
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
