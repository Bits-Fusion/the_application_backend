package repositories

import (
	"fmt"

	"github.com/Bits-Fusion/the_application_backend/database"
	"github.com/Bits-Fusion/the_application_backend/features/users/entities"
	"github.com/labstack/gommon/log"
)

type UserRepositoryImpl struct {
	db database.Database
}

func NewCockroachPostgresRepository(db database.Database) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) InsertUserData(in *entities.InsertUserDTO) error {

	data := &entities.User{
		Username:    in.Username,
		Role:        entities.Role(in.Role),
		Email:       in.Email,
		PhoneNumber: in.PhoneNumber,
		Password:    in.Password,
		FirstName:   in.FirstName,
		LastName:    in.LastName,
		IsActive:    true,
	}

	result := r.db.GetDb().Create(data)

	if result.Error != nil {
		log.Errorf("InsertCockroachData: %v", result.Error)
		return result.Error
	}

	log.Debugf("InsertCockroachData: %v", result.RowsAffected)
	return nil
}

func (r *UserRepositoryImpl) GetUserData(filterBy entities.FilterField, values ...string) (entities.User, error) {
	var user entities.User
	db := r.db.GetDb()

	switch filterBy {
	case entities.FilterByID:
		return r.querySingleField("id = ?", values, &user)
	case entities.FilterByUsername:
		return r.querySingleField("username = ?", values, &user)
	case entities.FilterByEmail:
		return r.querySingleField("email = ?", values, &user)
	case entities.FilterByPhoneNumber:
		return r.querySingleField("phone_number = ?", values, &user)
	case entities.FilterByAll:
		if len(values) < 3 {
			return user, fmt.Errorf("missing query value")
		}
		err := db.Where(
			"username = ? OR email = ? OR phone_number = ?",
			values[0], values[1], values[2],
		).First(&user).Error
		return user, err
	default:
		return user, fmt.Errorf("unsupported filter: %s", filterBy)
	}
}

func (r *UserRepositoryImpl) querySingleField(query string, values []string, user *entities.User) (entities.User, error) {
	if len(values) < 1 {
		return *user, fmt.Errorf("missing query value")
	}
	err := r.db.GetDb().Where(query, values[0]).First(user).Error
	return *user, err
}
