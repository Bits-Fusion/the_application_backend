package repositories

import "github.com/Bits-Fusion/the_application_backend/features/leads/entities"

type LeadRepository interface {
	CreateLead(in *entities.InsertLead) error
}
