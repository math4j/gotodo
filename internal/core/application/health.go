package application

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type Health struct {
}

func NewHealth() *Health {

	return &Health{}
}

func (h *Health) GetHealth(e echo.Context) error {

	response := map[string]any{
		"code":      0,
		"status":    "UP",
		"timestamp": time.Now(),
	}

	return e.JSON(http.StatusOK, response)
}
