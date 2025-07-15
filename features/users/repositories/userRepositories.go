package repositories

import "github.com/Bits-Fusion/the_application_backend/features/users/entities"

type UserRepository interface {
	InsertUserData(in *entities.InsertUserDTO) error
	GetUserData(filterBy entities.FilterField, values ...string) (entities.User, error)
	ListUsers(params entities.FilterParams) ([]entities.User, error)
}
