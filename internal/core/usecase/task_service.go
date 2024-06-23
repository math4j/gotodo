package usecase

import (
	"log/slog"

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

	slog.Info("a new task will be created")

	return s.repository.Create(task)
}

func (s *TaskService) FindById(id int64) (*model.Task, error) {

	slog.Info("getting new task", "taskId", id)
	return s.repository.FindById(id)
}

func (s *TaskService) FindAll() ([]*model.Task, error) {

	slog.Info("getting tasks")
	return s.repository.FindAll()
}

func (s *TaskService) DeleteById(id int64) error {

	slog.Info("delete new task", "taskId", id)
	return s.repository.DeleteById(id)
}
