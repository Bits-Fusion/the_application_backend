package usecases

import (
	"github.com/Bits-Fusion/the_application_backend/features/leads/entities"
	"github.com/Bits-Fusion/the_application_backend/features/leads/models"
	userModel "github.com/Bits-Fusion/the_application_backend/features/users/models"
)

type LeadUsecase interface {
	CreateLead(in *models.LeadInsertDTO) error
	ListLeads(params models.LeadFilterProps) ([]entities.Lead, error)
	UpdateLead(in *models.LeadUpdateDTO, leadId string) (entities.Lead, error)
	DeleteLead(mode userModel.DeleteMode, leadId ...string) (bool, error)
}
