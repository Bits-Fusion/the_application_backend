package usecases

import (
	"github.com/Bits-Fusion/the_application_backend/features/users/entities"
	"github.com/Bits-Fusion/the_application_backend/features/users/models"
)

type UserUsecase interface {
	CreateUser(in *models.UserModel) error
	UpdateUser(in *models.UserUpdateModel, userId string) (entities.User, error)
	ListUser(params entities.FilterParams) ([]entities.User, error)
	GetUserData(filterBy entities.FilterField, values ...string) (entities.User, error)
}
