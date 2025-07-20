package usecases

import (
	"github.com/Bits-Fusion/the_application_backend/features/tasks/entities"
	"github.com/Bits-Fusion/the_application_backend/features/tasks/models"
	userModel "github.com/Bits-Fusion/the_application_backend/features/users/models"
)

type TaskUsecase interface {
	CreateTask(in *models.TaskModel) error
	ListTask(filterOpts models.TaskFilterProps) ([]entities.Task, error)
	UpdateTask(in *models.TaskModelUpdate, taskId string) (entities.Task, error)
	DeleteTask(deleteMode userModel.DeleteMode, taskIds ...string) (bool, error)
}
