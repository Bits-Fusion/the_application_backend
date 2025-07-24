package repositories

import (
	"errors"
	"fmt"
	"github.com/Bits-Fusion/the_application_backend/database"
	"github.com/Bits-Fusion/the_application_backend/features/tasks/entities"
	"github.com/Bits-Fusion/the_application_backend/features/tasks/models"
	userEntity "github.com/Bits-Fusion/the_application_backend/features/users/entities"
	userModel "github.com/Bits-Fusion/the_application_backend/features/users/models"
	"github.com/google/uuid"
)

type taskRepository struct {
	db database.Database
}

func NewTaskRepository(db database.Database) *taskRepository {
	return &taskRepository{
		db: db,
	}
}

func (r *taskRepository) CreateTask(in *entities.InsertTask) error {
	var assignedUsers []userEntity.User

	for _, userID := range in.AssignedEmployees {
		parsedID, err := uuid.Parse(userID)
		if err != nil {
			return fmt.Errorf("invalid user ID %s: %w", userID, err)
		}
		assignedUsers = append(assignedUsers, userEntity.User{Id: parsedID, Role: "user"})
	}

	task := &entities.Task{
		Title:             in.Title,
		Description:       in.Description,
		Date:              in.Date,
		Place:             in.Place,
		Deadline:          in.Deadline,
		Priority:          in.Priority,
		Status:            in.Status,
		AssignedEmployees: assignedUsers,
	}

	return r.db.GetDb().Debug().Create(task).Error
}

func (r *taskRepository) ListTask(params models.TaskFilterProps) ([]entities.Task, error) {
	var tasks []entities.Task

	page := max(params.Page, 1)
	limit := params.Limit
	if limit <= 0 {
		limit = 10
	}
	offset := (page - 1) * limit

	order := "id desc"
	if params.OrderBy != "" {
		order = params.OrderBy
	}

	query := r.db.GetDb().Model(&entities.Task{}).Preload("AssignedEmployees")

	if params.AssignedTo != uuid.Nil {
		query = query.
			Joins("JOIN task_assignees ON task_assignees.task_id = tasks.id").
			Where("task_assignees.user_id = ?", params.AssignedTo)
	}

	if params.Priority != models.AllPriority && params.Priority != "" {
		query = query.Where("priority = ?", string(params.Priority))
	}

	if params.Status != models.AllStatus && params.Status != "" {
		query = query.Where("status = ?", string(params.Status))
	}

	err := query.Order(order).Limit(int(limit)).Offset(int(offset)).Find(&tasks).Error
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *taskRepository) UpdateTask(in *entities.UpdateTask, taskId string) (entities.Task, error) {

	var task entities.Task

	if err := r.db.GetDb().First(&task, "id = ?", taskId).Error; err != nil {
		return entities.Task{}, err
	}

	if in.Title != nil {
		task.Title = *in.Title
	}

	if in.Description != nil {
		task.Description = *in.Description
	}

	if in.Date != nil {
		task.Date = *in.Date
	}

	if in.Place != nil {
		task.Place = *in.Place
	}

	if in.Deadline != nil {
		task.Deadline = *in.Deadline
	}

	if in.Priority != nil {
		task.Priority = *in.Priority
	}

	if in.Status != nil {
		task.Status = *in.Status
	}

	if in.AssignedEmployeeIDs != nil {
		var assignedUsers []userEntity.User

		for _, userID := range *in.AssignedEmployeeIDs {
			assignedUsers = append(assignedUsers, userEntity.User{Id: userID, Role: "user"})
		}
		if err := r.db.GetDb().Model(&task).Association("AssignedEmployees").Replace(assignedUsers); err != nil {
			return entities.Task{}, err
		}
	}

	if err := r.db.GetDb().Save(&task).Error; err != nil {
		return entities.Task{}, err
	}

	if err := r.db.GetDb().Preload("AssignedEmployees").First(&task, "id = ?", task.Id).Error; err != nil {
		return entities.Task{}, err
	}

	return task, nil
}

func (r *taskRepository) DeleteTask(deletionState userModel.DeleteMode, taskIds ...string) (bool, error) {
	switch deletionState {
	case userModel.Single:
		ctx := r.db.GetDb().Delete(&entities.Task{}, taskIds[0])
		if ctx.RowsAffected == 0 {
			return false, errors.New("no recored found with this Id")
		}
	case userModel.All:
		ctx := r.db.GetDb().Delete(&entities.Task{}, taskIds)
		if ctx.RowsAffected == 0 {
			return false, errors.New("no recored found with this Id")
		}
	default:
		return false, errors.New("invalid mode")
	}

	return true, nil
}
