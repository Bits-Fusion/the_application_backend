package repositories

import (
	"errors"
	"fmt"

	"github.com/Bits-Fusion/the_application_backend/database"
	"github.com/Bits-Fusion/the_application_backend/features/leads/entities"
	"github.com/Bits-Fusion/the_application_backend/features/leads/models"
	userEntity "github.com/Bits-Fusion/the_application_backend/features/users/entities"
	userModel "github.com/Bits-Fusion/the_application_backend/features/users/models"
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

func (r *leadRepository) ListLeads(params models.LeadFilterProps) ([]entities.Lead, error) {
	var leads []entities.Lead

	page := max(params.Page, 1)
	limit := params.Limit
	if limit <= 0 {
		limit = 10
	}
	offset := (page - 1) * limit

	order := "id desc"
	if params.OrderBy != "" {
		order = params.OrderBy
	}

	query := r.db.GetDb().Model(&entities.Lead{}).Preload("AssignedEmployees")

	if params.AssignedTo != uuid.Nil {
		query = query.
			Joins("JOIN lead_assignees ON lead_assignees.lead_id = leads.id").
			Where("lead_assignees.user_id = ?", params.AssignedTo)
	}

	if params.Priority != "" {
		query = query.Where("priority = ?", string(params.Priority))
	}

	if params.Stage != "" {
		query = query.Where("stage = ?", params.Stage)
	}

	err := query.Order(order).Limit(int(limit)).Offset(int(offset)).Find(&leads).Error
	if err != nil {
		return nil, err
	}

	return leads, nil
}

func (r *leadRepository) UpdateLead(in *entities.InsertLead, leadId string) (entities.Lead, error) {

	var lead entities.Lead

	if err := r.db.GetDb().First(&lead, "id = ?", leadId).Error; err != nil {
		return entities.Lead{}, err
	}

	if in.Name != "" {
		lead.Name = in.Name
	}

	if in.Email != "" {
		lead.Email = in.Email
	}

	if in.Address != "" {
		lead.Address = in.Address
	}

	if in.PhoneNumber != "" {
		lead.PhoneNumber = in.PhoneNumber
	}

	if in.Stage != "" {
		lead.Stage = in.Stage
	}

	if in.Details != "" {
		lead.Details = in.Details
	}

	if in.LeadValue != 0 {
		lead.LeadValue = in.LeadValue
	}

	if !in.MeetingDate.IsZero() {
		lead.MeetingDate = in.MeetingDate
	}

	if in.ContactPerson != "" {
		lead.ContactPerson = in.ContactPerson
	}

	if in.Company != "" {
		lead.Company = ""
	}

	if in.Priority != "" {
		lead.Priority = in.Priority
	}

	if in.AssignedEmployees != nil {
		var assignedUsers []userEntity.User

		for _, userID := range in.AssignedEmployees {
			parsedID, err := uuid.Parse(userID)
			if err != nil {
				return entities.Lead{}, fmt.Errorf("invalid user ID %s: %w", userID, err)
			}
			assignedUsers = append(assignedUsers, userEntity.User{Id: parsedID, Role: "user"})
		}
		lead.AssignedEmployees = assignedUsers
	}

	if err := r.db.GetDb().Save(&lead).Error; err != nil {
		return entities.Lead{}, err
	}

	if err := r.db.GetDb().Preload("AssignedEmployees").First(&lead, "id = ?", lead.Id).Error; err != nil {
		return entities.Lead{}, err
	}

	return lead, nil
}

func (r *leadRepository) DeleteLead(deletionState userModel.DeleteMode, leadId ...string) (bool, error) {
	switch deletionState {
	case userModel.Single:
		ctx := r.db.GetDb().Delete(&entities.Lead{}, leadId[0])
		if ctx.RowsAffected == 0 {
			return false, errors.New("no recored found with this Id")
		}
	case userModel.All:
		ctx := r.db.GetDb().Delete(&entities.Lead{}, leadId)
		if ctx.RowsAffected == 0 {
			return false, errors.New("no recored found with this Id")
		}
	default:
		return false, errors.New("invalid mode")
	}
	return true, nil
}
