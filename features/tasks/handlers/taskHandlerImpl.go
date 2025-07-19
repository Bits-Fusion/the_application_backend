package handlers

import (
	"github.com/Bits-Fusion/the_application_backend/features/tasks/usecases"
	"github.com/labstack/echo/v4"
)

type taskHandlerImpl struct {
	taskUsecase usecases.TaskUsecase
}

func (h *taskHandlerImpl) CreateTask(c echo.Context) error {
	return nil
}
