package memory

import (
	"errors"

	"github.com/math4j/gotodo/internal/infra/model"
)

type TaskRepository struct {
	tasks []*model.Task
}

func NewTaskRepository() *TaskRepository {

	return &TaskRepository{
		tasks: make([]*model.Task, 0),
	}
}

func (r *TaskRepository) Create(task *model.Task) error {

	r.tasks = append(r.tasks, task)
	return nil
}

func (r *TaskRepository) FindById(id int64) (*model.Task, error) {

	for _, task := range r.tasks {

		if task.Id == id {
			return task, nil
		}
	}

	return nil, errors.New("task not found")
}

func (r *TaskRepository) FindAll() ([]*model.Task, error) {
	return r.tasks, nil
}

func (r *TaskRepository) DeleteById(id int64) error {

	for index, task := range r.tasks {

		if task.Id != id {
			r.tasks = append(r.tasks[:index], r.tasks[index+1:]...)
			return nil
		}
	}

	return nil
}
