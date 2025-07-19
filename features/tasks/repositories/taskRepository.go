package repositories

import "github.com/Bits-Fusion/the_application_backend/features/tasks/entities"

type TaskRepository interface {
	CreateTask(in *entities.InsertTask) error
}
