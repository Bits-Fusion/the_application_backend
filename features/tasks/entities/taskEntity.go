package entities

import (
	"time"

	userEntity "github.com/Bits-Fusion/the_application_backend/features/users/entities"
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
	Id                 int32           `json:"id" gorm:"primeryKey"`
	Title              string          `json:"title"`
	Description        string          `json:"description"`
	Date               time.Time       `json:"date"`
	Place              string          `json:"place"`
	Deadline           time.Time       `json:"deadline"`
	AssignedEmployeeID string          `json:"assigned_employee_id"`
	AssignedEmployee   userEntity.User `json:"assigned_employee" gorm:"foreignKey:AssignedEmployeeID;references:Id;OnDelete:SET NULL"`
	Priority           TaskPriority    `json:"priority" gorm:"type:priority_enum"`
	Status             Status          `json:"status" gorm:"type:status_enum"`
	UpdatedAt          time.Time       `json:"updated_at"`
	CreatedAt          time.Time       `json:"created_at"`
	// assigned_client varchar(255)
}

type InsertTask struct {
	Title              string
	Description        string
	Date               time.Time
	Place              string
	Deadline           time.Time
	AssignedEmployeeID string
	Priority           TaskPriority
	Status             Status
}
