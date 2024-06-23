package usecase

import (
	"github.com/math4j/gotodo/internal/infra/model"
	"github.com/math4j/gotodo/internal/infra/repository"
)

type TaskService struct {
	repository repository.TaskRepository
}

func NewTaskService(repository repository.TaskRepository) *TaskService {

	return &TaskService{
		repository: repository,
	}
}

func (s *TaskService) Create(task *model.Task) error {

	return s.repository.Create(task)
}

func (s *TaskService) FindById(id int64) (*model.Task, error) {

	return s.repository.FindById(id)
}

func (s *TaskService) DeleteById(id int64) error {

	return s.repository.DeleteById(id)
}
