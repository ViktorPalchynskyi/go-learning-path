package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/username/lesson_11/internal/domain"
	"github.com/username/lesson_11/internal/service"
)

type mockRepo struct {
	tasks map[string]*domain.Task
}

func newMockRepo() *mockRepo {
	return &mockRepo{tasks: make(map[string]*domain.Task)}
}

func (r *mockRepo) Save(task *domain.Task) error {
	r.tasks[task.ID] = task
	return nil
}

func (r *mockRepo) FindById(id string) (*domain.Task, error) {
	task, ok := r.tasks[id]
	if !ok {
		return nil, domain.ErrNotFound
	}
	return task, nil
}

func (r *mockRepo) FindAll() ([]*domain.Task, error) {
	result := make([]*domain.Task, 0, len(r.tasks))
	for _, t := range r.tasks {
		result = append(result, t)
	}
	return result, nil
}

func (r *mockRepo) Update(task *domain.Task) error {
	if _, ok := r.tasks[task.ID]; !ok {
		return domain.ErrNotFound
	}
	r.tasks[task.ID] = task
	return nil
}

func (r *mockRepo) Delete(id string) error {
	if _, ok := r.tasks[id]; !ok {
		return domain.ErrNotFound
	}
	delete(r.tasks, id)
	return nil
}

func withURLParam(r *http.Request, key, value string) *http.Request {
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add(key, value)
	return r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
}

func TestGetTask_Success(t *testing.T) {
	repo := newMockRepo()
	repo.tasks["1"] = &domain.Task{ID: "1", Title: "Test"}
	h := NewHandler(service.NewTaskService(repo))

	req := httptest.NewRequest("GET", "/tasks/1", nil)
	req = withURLParam(req, "id", "1")
	w := httptest.NewRecorder()

	h.GetTask(w, req)

	if w.Code != 200 {
		t.Errorf("status = %d, want 200", w.Code)
	}

	var task domain.Task
	json.NewDecoder(w.Body).Decode(&task)
	if task.ID != "1" {
		t.Errorf("ID = %q, want %q", task.ID, "1")
	}
}

func TestGetTask_NotFound(t *testing.T) {
	h := NewHandler(service.NewTaskService(newMockRepo()))

	req := httptest.NewRequest("GET", "/tasks/999", nil)
	req = withURLParam(req, "id", "999")
	w := httptest.NewRecorder()

	h.GetTask(w, req)

	if w.Code != 404 {
		t.Errorf("status = %d, want 404", w.Code)
	}
}

func TestCreateTask_Success(t *testing.T) {
	h := NewHandler(service.NewTaskService(newMockRepo()))

	body, _ := json.Marshal(CreateTaskRequest{Title: "New Task"})
	req := httptest.NewRequest("POST", "/tasks", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.CreateTask(w, req)

	if w.Code != 201 {
		t.Errorf("status = %d, want 201", w.Code)
	}

	var task domain.Task
	json.NewDecoder(w.Body).Decode(&task)
	if task.Title != "New Task" {
		t.Errorf("Title = %q, want %q", task.Title, "New Task")
	}
}

func TestCreateTask_BadRequest(t *testing.T) {
	h := NewHandler(service.NewTaskService(newMockRepo()))

	req := httptest.NewRequest("POST", "/tasks", bytes.NewReader([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.CreateTask(w, req)

	if w.Code != 400 {
		t.Errorf("status = %d, want 400", w.Code)
	}
}
