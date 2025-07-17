package repositories

import (
	"github.com/Bits-Fusion/the_application_backend/features/users/entities"
	"github.com/Bits-Fusion/the_application_backend/features/users/models"
)

type UserRepository interface {
	InsertUserData(in *entities.InsertUserDTO) error
	GetUserData(filterBy entities.FilterField, values ...string) (entities.User, error)
	ListUsers(params entities.FilterParams) ([]entities.User, error)
	UpdateUser(in *entities.InsertUserDTO, userId string) (entities.User, error)
	DeleteUser(deleteMode models.DeleteMode, userId ...string) (bool, error)
}
