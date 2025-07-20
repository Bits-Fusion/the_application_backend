package models

import (
	"time"

	"github.com/Bits-Fusion/the_application_backend/features/tasks/entities"
	"github.com/google/uuid"
)

type TaskModel struct {
	Title              string    `validate:"required" json:"title"`
	Description        string    `validate:"required" json:"description"`
	Date               time.Time `validate:"required" json:"date"`
	Place              string    `validate:"required" json:"place"`
	Deadline           time.Time `validate:"required" json:"deadline"`
	AssignedEmployeeID string    `validate:"required" json:"assignedTo"`
	Priority           string    `validate:"required,oneof=low medium high" json:"priority"`
	Status             string    `validate:"required,oneof=completed inprogress canceled" json:"status"`
}

type TaskModelUpdate struct {
	Title              string                `json:"title"`
	Description        string                `json:"description"`
	Date               time.Time             `json:"date"`
	Place              string                `json:"place"`
	Deadline           time.Time             `json:"deadline"`
	AssignedEmployeeID string                `json:"assignedTo"`
	Priority           entities.TaskPriority `validate:"omitempty,oneof=low medium high" json:"priority"`
	Status             entities.Status       `validate:"omitempty,oneof=completed inprogress canceled" json:"status"`
}

type PriorityFilterOpt string

const (
	High        PriorityFilterOpt = "high"
	Mid         PriorityFilterOpt = "mid"
	Low         PriorityFilterOpt = "low"
	AllPriority PriorityFilterOpt = "all"
)

type StatusFiterOpt string

const (
	Complete   StatusFiterOpt = "complete"
	InProgress StatusFiterOpt = "inprogress"
	Canceled   StatusFiterOpt = "canceled"
	AllStatus  StatusFiterOpt = "all"
)

type TaskFilterProps struct {
	AssignedTo uuid.UUID
	Limit      int32
	Page       int32
	OrderBy    string
	Priority   PriorityFilterOpt
	Status     StatusFiterOpt
}
