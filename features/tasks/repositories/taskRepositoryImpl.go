package repositories

import (
	"github.com/Bits-Fusion/the_application_backend/database"
	"github.com/Bits-Fusion/the_application_backend/features/tasks/entities"
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
