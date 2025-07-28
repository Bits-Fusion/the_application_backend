package entities

import (
	"time"

	userEntity "github.com/Bits-Fusion/the_application_backend/features/users/entities"
)

type LeadsPriority string

const (
	High LeadsPriority = "high"
	Mid  LeadsPriority = "mid"
	Low  LeadsPriority = "low"
)

type Lead struct {
	Id                int32             `json:"id" gorm:"primaryKey"`
	Name              string            `json:"name"`
	ContactPerson     string            `json:"contact_person"`
	Email             string            `json:"email"`
	PhoneNumber       string            `json:"phone_number"`
	Company           string            `json:"company"`
	Address           string            `json:"address"`
	Stage             string            `json:"stage"`
	MeetingDate       time.Time         `json:"meeting_date"`
	Details           string            `json:"details"`
	AssignedEmployees []userEntity.User `json:"assigned_employees" gorm:"many2many:lead_assignees;constraint:OnDelete:CASCADE;"`
	Priority          LeadsPriority     `json:"priority" gorm:"type:priority_enum"`
	CreatedAt         time.Time         `json:"created_at"`
	UpdatedAT         time.Time         `json:"updated_at"`
	LeadValue         int32             `json:"lead_value"`
}

type InsertLead struct {
	Name              string        `json:"name"`
	ContactPerson     string        `json:"contact_person"`
	Email             string        `json:"email"`
	PhoneNumber       string        `json:"phone_number"`
	Company           string        `json:"company"`
	Address           string        `json:"address"`
	Stage             string        `json:"stage"`
	MeetingDate       time.Time     `json:"meeting_date"`
	Details           string        `json:"details"`
	Priority          LeadsPriority `json:"priority" `
	LeadValue         int32         `json:"lead_value"`
	AssignedEmployees []string
}
