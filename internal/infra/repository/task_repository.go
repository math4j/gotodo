package repository

import "github.com/math4j/gotodo/internal/infra/model"

type TaskRepository interface {
	Create(*model.Task) error
	FindById(int64) (*model.Task, error)
	DeleteById(int64) error
}
