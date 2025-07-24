package usecases

import "github.com/Bits-Fusion/the_application_backend/features/leads/models"

type LeadUsecase interface {
	CreateLead(in *models.LeadInsertDTO) error
}
