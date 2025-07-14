package usecases

import (
	"github.com/Bits-Fusion/the_application_backend/features/users/models"
)

type UserUsecase interface {
	CreateUser(in *models.UserModel) error
}
