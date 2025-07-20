package usecases

import (
	"github.com/Bits-Fusion/the_application_backend/features/tasks/entities"
	"github.com/Bits-Fusion/the_application_backend/features/tasks/models"
	"github.com/Bits-Fusion/the_application_backend/features/tasks/repositories"
)

type taskUsecaseImpl struct {
	taskRepo repositories.TaskRepository
}

func NewTaskUsecase(taskRepo repositories.TaskRepository) *taskUsecaseImpl {
	return &taskUsecaseImpl{
		taskRepo: taskRepo,
	}
}

func (u *taskUsecaseImpl) CreateTask(in *models.TaskModel) error {
	task := entities.InsertTask{
		Title:              in.Title,
		Description:        in.Description,
		Date:               in.Date,
		Place:              in.Place,
		Deadline:           in.Deadline,
		AssignedEmployeeID: in.AssignedEmployeeID,
		Priority:           entities.TaskPriority(in.Priority),
		Status:             entities.Status(in.Status),
	}
	return u.taskRepo.CreateTask(&task)
}

func (u *taskUsecaseImpl) ListTask(filterOpts models.TaskFilterProps) ([]entities.Task, error) {
	return u.taskRepo.ListTask(filterOpts)
}
