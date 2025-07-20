package handlers

import (
	"github.com/labstack/echo/v4"
)

type TaskHandler interface {
	CreateTask(c echo.Context) error
	ListTasks(c echo.Context) error
	UpdateTask(c echo.Context) error
	DeleteTask(c echo.Context) error
}
