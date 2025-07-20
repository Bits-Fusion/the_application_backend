package repositories

import (
	"github.com/Bits-Fusion/the_application_backend/features/tasks/entities"
	"github.com/Bits-Fusion/the_application_backend/features/tasks/models"
	userModel "github.com/Bits-Fusion/the_application_backend/features/users/models"
)

type TaskRepository interface {
	CreateTask(in *entities.InsertTask) error
	ListTask(filterOpts models.TaskFilterProps) ([]entities.Task, error)
	UpdateTask(in *entities.UpdateTask, taskId string) (entities.Task, error)
	DeleteTask(deletionState userModel.DeleteMode, taskIds ...string) (bool, error)
}
