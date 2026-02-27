package service

import "github.com/username/pomodoro/internal/ports"

type TaskService struct {
    repo ports.TaskRepository
}