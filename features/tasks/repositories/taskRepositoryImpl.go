package repositories

import (
	"github.com/Bits-Fusion/the_application_backend/database"
	"github.com/Bits-Fusion/the_application_backend/features/tasks/entities"
	"github.com/Bits-Fusion/the_application_backend/features/tasks/models"
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
	task := &entities.Task{
		Title:              in.Title,
		Description:        in.Description,
		Date:               in.Date,
		Place:              in.Place,
		Deadline:           in.Deadline,
		AssignedEmployeeID: in.AssignedEmployeeID,
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

	// AssignedTo (UUID as string)
	if params.AssignedTo != uuid.Nil {
		query = query.Where("assigned_employee_id = ?", params.AssignedTo.String())
	}

	// Priority filter
	if params.Priority != models.AllPriority && params.Priority != "" {
		query = query.Where("priority = ?", string(params.Priority))
	}

	// Status filter
	if params.Status != models.AllStatus && params.Status != "" {
		query = query.Where("status = ?", string(params.Status))
	}

	err := query.Order(order).Limit(int(limit)).Offset(int(offset)).Find(&tasks).Error
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
