package repositories

import (
	"fmt"

	"github.com/Bits-Fusion/the_application_backend/database"
	"github.com/Bits-Fusion/the_application_backend/features/leads/entities"
	userEntity "github.com/Bits-Fusion/the_application_backend/features/users/entities"
	"github.com/google/uuid"
)

type leadRepository struct {
	db database.Database
}

func NewLeadRepository(db database.Database) *leadRepository {
	return &leadRepository{
		db: db,
	}
}

func (r *leadRepository) CreateLead(in *entities.InsertLead) error {

	var assignedUsers []userEntity.User

	for _, userID := range in.AssignedEmployees {
		parsedID, err := uuid.Parse(userID)
		if err != nil {
			return fmt.Errorf("invalid user ID %s: %w", userID, err)
		}
		assignedUsers = append(assignedUsers, userEntity.User{Id: parsedID, Role: "user"})
	}

	data := &entities.Lead{
		Name:              in.Name,
		ContactPerson:     in.ContactPerson,
		Email:             in.Email,
		PhoneNumber:       in.PhoneNumber,
		Company:           in.Company,
		Address:           in.Address,
		Stage:             in.Stage,
		MeetingDate:       in.MeetingDate,
		Details:           in.Details,
		Priority:          in.Priority,
		LeadValue:         in.LeadValue,
		AssignedEmployees: assignedUsers,
	}

	if err := r.db.GetDb().Create(data).Error; err != nil {
		return err
	}

	return nil
}
