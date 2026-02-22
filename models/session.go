package models

import "time"

type SessionType string

const (
	WorkSession SessionType = "work"
	BreakSession SessionType = "break"
)

type Session struct {
	BaseEntity
	TaskID string
	StartTime time.Time
	EndTime time.Time
	Duration time.Duration
	Type SessionType 
}

func NewWorkSession(id, taskId string, duration time.Duration, startDate time.Time) *Session{
	return &Session{
		TaskID: taskId,
		Duration: duration,
		StartTime: startDate,
	}
}