package usecases

import (
	"github.com/Bits-Fusion/the_application_backend/features/tasks/entities"
	"github.com/Bits-Fusion/the_application_backend/features/tasks/models"
)

type TaskUsecase interface {
	CreateTask(in *models.TaskModel) error
	ListTask(filterOpts models.TaskFilterProps) ([]entities.Task, error)
}
