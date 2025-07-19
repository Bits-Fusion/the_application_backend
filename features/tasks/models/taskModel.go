package models

import (
	"time"
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
