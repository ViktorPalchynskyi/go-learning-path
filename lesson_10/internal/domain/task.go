package domain

type Task struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func NewTask(id string, title string) *Task{
	return &Task{
		ID: id,
		Title: title,
		Completed: false,
	}
}