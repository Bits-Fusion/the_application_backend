package handlers

import (
	"github.com/labstack/echo/v4"
)

type TaskHandler interface {
	CreateTask(c echo.Context) error
}
