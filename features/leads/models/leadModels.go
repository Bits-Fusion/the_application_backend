package models

import (
	"time"

	"github.com/google/uuid"
)

type LeadInsertDTO struct {
	Name                string    `validate:"required" json:"name"`
	ContactPerson       string    `validate:"required" json:"contactPerson"`
	Email               string    `validate:"required,email" json:"email"`
	PhoneNumber         string    `validate:"required" json:"phoneNumber"`
	Company             string    `validate:"required" json:"company"`
	Address             string    `validate:"required" json:"address"`
	Stage               string    `validate:"required" json:"stage"`
	MeetingDate         time.Time `validate:"required" json:"meetingDate"`
	Details             string    `json:"details"`
	Priority            string    `validate:"required,oneof=low medium high" json:"priority" `
	LeadValue           int32     `validate:"required" json:"leadValue"`
	AssignedEmployeeIDs []string  `validate:"omitempty,dive,required,uuid4" json:"assignedTo"`
}

type LeadUpdateDTO struct {
	Name                string    `validate:"omitempty" json:"name"`
	ContactPerson       string    `validate:"omitempty" json:"contactPerson"`
	Email               string    `validate:"omitempty,email" json:"email"`
	PhoneNumber         string    `validate:"omitempty" json:"phoneNumber"`
	Company             string    `validate:"omitempty" json:"company"`
	Address             string    `validate:"omitempty" json:"address"`
	Stage               string    `validate:"omitempty" json:"stage"`
	MeetingDate         time.Time `validate:"omitempty" json:"meetingDate"`
	Details             string    `json:"details"`
	Priority            string    `validate:"omitempty,oneof=low medium high" json:"priority" `
	LeadValue           int32     `validate:"omitempty" json:"leadValue"`
	AssignedEmployeeIDs []string  `validate:"omitempty,dive,required,uuid4" json:"assignedTo"`
}

type LeadFilterProps struct {
	Priority   string
	Stage      string
	AssignedTo uuid.UUID
	Limit      int32
	Page       int32
	OrderBy    string
}
