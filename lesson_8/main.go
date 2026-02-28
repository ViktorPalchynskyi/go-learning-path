package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func main()  {
	fmt.Println("Lesson 8")

	task := &Task{
		ID: "123",
		Title: "Test1",
	}
	data, err := json.Marshal(task)
	if err != nil {
		fmt.Errorf("%w", err)
	}
	fmt.Println(string(data))
}

type Task struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Description string `json:"description,omitempty"`
	Completed bool `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateTaskRequest struct {
	Title string `json:"title"`
	Description string `json:"description,omitempty"`
}

type TaskResponse struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
	CreatedAt string `json:"created_at"`
}

func toTaskReponse(t *Task) TaskResponse {
	return TaskResponse{
		ID: t.ID,
		Title: t.Title,
		Completed: t.Completed,
		CreatedAt: t.CreatedAt.Format(time.RFC3339),
	}
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

type TaskService interface {
	CreateTask(title string) (*Task, error)
}

type Handler struct {
	service TaskService
}

func NewHandler(service TaskService) *Handler {
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
	writeJSON(w, 201, toTaskReponse(task))
}