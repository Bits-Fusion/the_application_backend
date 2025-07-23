package entities

import (
	"time"

	userEntity "github.com/Bits-Fusion/the_application_backend/features/users/entities"
	"github.com/google/uuid"
)

type TaskPriority string

const (
	High TaskPriority = "high"
	Mid  TaskPriority = "mid"
	Low  TaskPriority = "low"
)

type Status string

const (
	Complete   Status = "complete"
	InProgress Status = "inprogress"
	Canceled   Status = "canceled"
)

type Task struct {
	Id                int32             `json:"id" gorm:"primeryKey"`
	Title             string            `json:"title"`
	Description       string            `json:"description"`
	Date              time.Time         `json:"date"`
	Place             string            `json:"place"`
	Deadline          time.Time         `json:"deadline"`
	AssignedEmployees []userEntity.User `json:"assigned_employees" gorm:"many2many:task_assignees;"`
	Priority          TaskPriority      `json:"priority" gorm:"type:priority_enum"`
	Status            Status            `json:"status" gorm:"type:status_enum"`
	UpdatedAt         time.Time         `json:"updated_at"`
	CreatedAt         time.Time         `json:"created_at"`
	// assigned_client varchar(255)
}

type InsertTask struct {
	Title             string
	Description       string
	Date              time.Time
	Place             string
	Deadline          time.Time
	AssignedEmployees []string
	Priority          TaskPriority
	Status            Status
}

type UpdateTask struct {
	Title               *string
	Description         *string
	Date                *time.Time
	Place               *string
	Deadline            *time.Time
	AssignedEmployeeIDs *[]uuid.UUID
	Priority            *TaskPriority
	Status              *Status
}
