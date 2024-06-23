package mysql

import (
	"database/sql"
	"log/slog"

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

	stmt, err := r.db.Prepare("INSERT INTO tasks(name, description) VALUES(?, ?)")

	if err != nil {
		slog.Error("could not prepare statement", "error", err)
		return err
	}

	_, err = stmt.Exec(task.Name, task.Description)

	if err != nil {

		return err
	}

	slog.Info("task created successfully")
	return nil
}

func (r *TaskRepository) FindById(id int64) (*model.Task, error) {

	row := r.db.QueryRow("SELECT * FROM tasks WHERE id = ?", id)
	task := &model.Task{}
	err := row.Scan(&task.Id, &task.Name, &task.Description)

	if err != nil {
		return nil, err
	}

	return task, nil
}

func (r *TaskRepository) FindAll() ([]*model.Task, error) {

	rows, err := r.db.Query("SELECT * FROM tasks")

	if err != nil {
		return nil, err
	}

	tasks := []*model.Task{}

	for rows.Next() {
		task := &model.Task{}
		err = rows.Scan(&task.Id, &task.Name, &task.Description)

		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r *TaskRepository) DeleteById(id int64) error {

	stmt, err := r.db.Prepare("DELETE FROM tasks WHERE id = ?")

	if err != nil {
		slog.Error("could not prepare delete sql statement", "err", err.Error())
		return err
	}

	_, err = stmt.Exec(id)

	if err != nil {
		slog.Error("could not execute sql query", "err", err.Error())
		return err
	}

	return nil
}
