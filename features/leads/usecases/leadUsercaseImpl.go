package usecases

import (
	"github.com/Bits-Fusion/the_application_backend/features/leads/entities"
	"github.com/Bits-Fusion/the_application_backend/features/leads/models"
	"github.com/Bits-Fusion/the_application_backend/features/leads/repositories"
	userModel "github.com/Bits-Fusion/the_application_backend/features/users/models"
	userUsecase "github.com/Bits-Fusion/the_application_backend/features/users/usecases"
)

type leadUsecase struct {
	leadRepo repositories.LeadRepository
}

func NewLeadUsecase(leadRepo repositories.LeadRepository) *leadUsecase {
	return &leadUsecase{
		leadRepo: leadRepo,
	}
}

func (u *leadUsecase) CreateLead(in *models.LeadInsertDTO) error {
	phone, err := userUsecase.StandardizePhoneNumber(in.PhoneNumber)

	if err != nil {
		return err
	}

	data := entities.InsertLead{
		Name:              in.Name,
		ContactPerson:     in.ContactPerson,
		Email:             in.Email,
		PhoneNumber:       phone,
		Company:           in.Company,
		Address:           in.Address,
		Stage:             in.Stage,
		MeetingDate:       in.MeetingDate,
		Details:           in.Details,
		Priority:          entities.LeadsPriority(in.Priority),
		LeadValue:         in.LeadValue,
		AssignedEmployees: in.AssignedEmployeeIDs,
	}
	return u.leadRepo.CreateLead(&data)
}

func (u *leadUsecase) ListLeads(params models.LeadFilterProps) ([]entities.Lead, error) {
	return u.leadRepo.ListLeads(params)
}

func (u *leadUsecase) UpdateLead(in *models.LeadUpdateDTO, leadId string) (entities.Lead, error) {

	data := entities.InsertLead{
		Name:              in.Name,
		ContactPerson:     in.ContactPerson,
		Email:             in.Email,
		Company:           in.Company,
		Address:           in.Address,
		Stage:             in.Stage,
		MeetingDate:       in.MeetingDate,
		Details:           in.Details,
		Priority:          entities.LeadsPriority(in.Priority),
		LeadValue:         in.LeadValue,
		AssignedEmployees: in.AssignedEmployeeIDs,
	}

	if in.PhoneNumber != "" {
		phone, err := userUsecase.StandardizePhoneNumber(in.PhoneNumber)
		if err != nil {
			return entities.Lead{}, err
		}
		data.PhoneNumber = phone
	}

	return u.leadRepo.UpdateLead(&data, leadId)
}

func (u *leadUsecase) DeleteLead(mode userModel.DeleteMode, leadId ...string) (bool, error) {
	return u.leadRepo.DeleteLead(mode, leadId...)
}
