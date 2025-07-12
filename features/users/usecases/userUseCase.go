package usecases

import "github.com/Bits-Fusion/the_application_backend/features/users/entities"

type UserUsecase interface {
	CreateUser(in *entities.InsertUserDTO)
}
