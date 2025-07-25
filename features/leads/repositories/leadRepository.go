package repositories

import (
	"github.com/Bits-Fusion/the_application_backend/features/leads/entities"
	"github.com/Bits-Fusion/the_application_backend/features/leads/models"
	userModel "github.com/Bits-Fusion/the_application_backend/features/users/models"
)

type LeadRepository interface {
	CreateLead(in *entities.InsertLead) error
	ListLeads(params models.LeadFilterProps) ([]entities.Lead, error)
	UpdateLead(in *entities.InsertLead, leadId string) (entities.Lead, error)
	DeleteLead(deletionState userModel.DeleteMode, leadId ...string) (bool, error)
}
