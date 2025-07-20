package repositories

import (
	"github.com/Bits-Fusion/the_application_backend/features/tasks/entities"
	"github.com/Bits-Fusion/the_application_backend/features/tasks/models"
)

type TaskRepository interface {
	CreateTask(in *entities.InsertTask) error
	ListTask(filterOpts models.TaskFilterProps) ([]entities.Task, error)
}
