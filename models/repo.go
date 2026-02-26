package models

type TaskRepository interface {
	Save(task *Task) error
	FindById(id string) (*Task, error)
	FindAll() ([]*Task, error)
	Delete(id string) error
}

type InMemoryTaskRepo struct {
	tasks map[string]*Task
}

func (repo *InMemoryTaskRepo) FindById(id string) (*Task, error) {
	task, ok := repo.tasks[id]
	if !ok {
		return nil, ErrTaskNotFound
	}

	return task, nil
}

func (repo *InMemoryTaskRepo) FindAll() ([]*Task, error) {
	result := make([]*Task, 0, len(repo.tasks))
	for _, t := range repo.tasks {
		result = append(result, t)
	}
	return result, nil
}

func (repo *InMemoryTaskRepo) Save(task *Task) error {
	repo.tasks[task.ID] = task
	return nil
}

func (repo *InMemoryTaskRepo) Delete(id string) error {

	if _, ok := repo.tasks[id]; !ok {
		return ErrTaskNotFound
	}

	delete(repo.tasks, id)
	return nil
}

func NewInMemoryTaskRepo() *InMemoryTaskRepo {
	return &InMemoryTaskRepo{tasks: make(map[string]*Task)}
}