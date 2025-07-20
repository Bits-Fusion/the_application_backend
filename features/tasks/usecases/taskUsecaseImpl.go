package usecases

import (
	"github.com/Bits-Fusion/the_application_backend/features/tasks/entities"
	"github.com/Bits-Fusion/the_application_backend/features/tasks/models"
	"github.com/Bits-Fusion/the_application_backend/features/tasks/repositories"
	userModel "github.com/Bits-Fusion/the_application_backend/features/users/models"
	"github.com/google/uuid"
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

func (u *taskUsecaseImpl) UpdateTask(in *models.TaskModelUpdate, taskId string) (entities.Task, error) {
	var updateTaskData entities.UpdateTask

	if in.Title != "" {
		updateTaskData.Title = &in.Title
	}

	if in.Description != "" {
		updateTaskData.Description = &in.Description
	}

	if !in.Date.IsZero() {
		updateTaskData.Date = &in.Date
	}

	if in.Place != "" {
		updateTaskData.Place = &in.Place
	}

	if !in.Deadline.IsZero() {
		updateTaskData.Deadline = &in.Deadline
	}

	if uid, err := uuid.Parse(in.AssignedEmployeeID); err != nil && uuid.Nil != uid {
		updateTaskData.AssignedEmployeeID = &uid
	}

	if in.Priority != "" {
		updateTaskData.Priority = &in.Priority
	}

	if in.Status != "" {
		updateTaskData.Status = &in.Status
	}

	if in.AssignedEmployeeID != "" {
		uid, err := uuid.Parse(in.AssignedEmployeeID)
		if err != nil || uid == uuid.Nil {
			return entities.Task{}, err
		}
		updateTaskData.AssignedEmployeeID = &uid
	}

	task, err := u.taskRepo.UpdateTask(&updateTaskData, taskId)

	if err != nil {
		return entities.Task{}, err
	}

	return task, nil
}

func (u *taskUsecaseImpl) DeleteTask(deleteMode userModel.DeleteMode, taskIds ...string) (bool, error) {
	return u.taskRepo.DeleteTask(deleteMode, taskIds...)
}
