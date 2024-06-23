package application

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/math4j/gotodo/internal/core/usecase"
	"github.com/math4j/gotodo/internal/infra/model"
	"github.com/math4j/gotodo/internal/infra/repository"
)

type Task struct {
	service *usecase.TaskService
}

func NewTask(repository repository.TaskRepository) *Task {

	service := usecase.NewTaskService(repository)
	return &Task{
		service: service,
	}
}

func (t *Task) Create(c echo.Context) error {

	task := &model.Task{}
	err := json.NewDecoder(c.Request().Body).Decode(task)

	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	err = t.service.Create(task)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, nil)
}

func (t *Task) Get(c echo.Context) error {

	id, err := strconv.ParseInt(c.QueryParam("taskId"), 10, 64)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	task, err := t.service.FindById(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, task)
}

func (t *Task) GetAll(c echo.Context) error {

	slog.Info("request received")
	tasks, err := t.service.FindAll()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, tasks)
}

func (t *Task) Delete(c echo.Context) error {

	id, err := strconv.ParseInt(c.QueryParam("taskId"), 10, 64)

	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	err = t.service.DeleteById(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, nil)
}
