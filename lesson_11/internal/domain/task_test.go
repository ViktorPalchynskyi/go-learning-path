package domain

import (
	"strings"
	"testing"
)

func TestNewTask(t *testing.T) {
	task := NewTask("1", "Learn Go")

	if task.ID != "1" {
		t.Errorf("ID = %q, want %q", task.ID, "1")
	}
	if task.Title != "Learn Go" {
		t.Errorf("Title = %q, want %q", task.Title, "Learn Go")
	}
	if task.Completed {
		t.Error("Completed = true, want false")
	}
}

func TestTask_Complete(t *testing.T) {
	task := NewTask("1", "Learn Go")
	task.Complete()

	if !task.Completed {
		t.Error("Completed = false, want true")
	}
}

func TestValidateTitle(t *testing.T) {
	tests := []struct {
		name    string
		title   string
		wantErr bool
	}{
		{"valid title", "Learn Go", false},
		{"empty title", "", true},
		{"spaces only", "   ", true},
		{"too long", strings.Repeat("a", 201), true},
		{"max length", strings.Repeat("a", 200), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateTitle(tt.title)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateTitle(%q) error = %v, wantErr %v", tt.title, err, tt.wantErr)
			}
		})
	}
}
