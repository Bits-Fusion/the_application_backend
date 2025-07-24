package models

import "time"

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
