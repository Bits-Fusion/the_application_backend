package repositories

import (
	"errors"
	"time"

	"github.com/Bits-Fusion/the_application_backend/database"
	"github.com/Bits-Fusion/the_application_backend/features/tasks/entities"
	"github.com/Bits-Fusion/the_application_backend/features/tasks/models"
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
	asUUID, err := uuid.Parse(in.AssignedEmployeeID)

	if err != nil {
		return err
	}

	task := &entities.Task{
		Title:              in.Title,
		Description:        in.Description,
		Date:               in.Date,
		Place:              in.Place,
		Deadline:           in.Deadline,
		AssignedEmployeeID: asUUID,
		Priority:           in.Priority,
		Status:             in.Status,
	}

	return r.db.GetDb().Create(task).Error
}

func (r *taskRepository) ListTask(params models.TaskFilterProps) ([]entities.Task, error) {
	var tasks []entities.Task

	page := max(params.Page, 1)
	limit := params.Limit
	if limit <= 0 {
		limit = 10
	}
	offset := (page - 1) * limit

	order := "id asc"
	if params.OrderBy != "" {
		order = params.OrderBy
	}

	query := r.db.GetDb().Model(&entities.Task{}).Preload("AssignedEmployee")

	if params.AssignedTo != uuid.Nil {
		query = query.Where("assigned_employee_id = ?", params.AssignedTo.String())
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

	if in.AssignedEmployeeID != nil {
		task.AssignedEmployeeID = *in.AssignedEmployeeID
	}

	if in.Priority != nil {
		task.Priority = *in.Priority
	}

	if in.Status != nil {
		task.Status = *in.Status
	}

	task.UpdatedAt = time.Now()

	if err := r.db.GetDb().Save(&task).Error; err != nil {
		return entities.Task{}, err
	}

	if err := r.db.GetDb().Preload("AssignedEmployee").First(&task, "id = ?", task.Id).Error; err != nil {
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
