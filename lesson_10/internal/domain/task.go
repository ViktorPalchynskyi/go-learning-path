package domain

type Task struct {
	ID        string
	Title     string
	Completed bool
}

func NewTask(id string, title string) *Task{
	return &Task{
		ID: id,
		Title: title,
		Completed: false,
	}
}