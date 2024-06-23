package memory

import (
	"database/sql"

	"github.com/math4j/gotodo/internal/infra/model"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {

	return &TaskRepository{
		db: db,
	}
}

func (r *TaskRepository) Create(task *model.Task) error {

	return nil
}

func (r *TaskRepository) FindById(id int64) (*model.Task, error) {

	return nil, nil
}

func (r *TaskRepository) DeleteById(id int64) (*model.Task, error) {

	return nil, nil
}
