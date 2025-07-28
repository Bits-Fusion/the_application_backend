package repositories

import (
	"errors"
	"fmt"
	"time"

	"github.com/Bits-Fusion/the_application_backend/database"
	"github.com/Bits-Fusion/the_application_backend/features/users/entities"
	"github.com/Bits-Fusion/the_application_backend/features/users/models"
	"github.com/labstack/gommon/log"
)

type userRepositoryImpl struct {
	db database.Database
}

func NewUserPostgresRepository(db database.Database) *userRepositoryImpl {
	return &userRepositoryImpl{db: db}
}

func (r *userRepositoryImpl) InsertUserData(in *entities.InsertUserDTO) error {

	data := &entities.User{
		Id:          in.Id,
		Username:    in.Username,
		Role:        entities.UserRole,
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

func (r *userRepositoryImpl) ListUsers(params entities.FilterParams) ([]entities.User, error) {
	var users []entities.User

	page := max(params.Page, 1)

	limit := params.Limit
	if limit <= 0 {
		limit = 10
	}

	offset := (page - 1) * limit
	order := "id asc"
	if params.OrderBy != "" {
		order = params.OrderBy
	}

	if err := r.db.GetDb().Order(order).Limit(int(limit)).Offset(int(offset)).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepositoryImpl) DeleteUser(deleteMode models.DeleteMode, userId ...string) (bool, error) {
	switch deleteMode {
	case models.All:
		ctx := r.db.GetDb().Delete(&entities.User{}, userId)
		if ctx.RowsAffected == 0 {
			return false, errors.New("no recored found with this Id")
		}
		return true, nil
	case models.Single:
		ctx := r.db.GetDb().Delete(&entities.User{}, userId)
		if ctx.RowsAffected == 0 {
			return false, errors.New("no recored found with this Id")
		}
		return true, nil
	default:
		return false, errors.New("invalid deletion mode")
	}
}

func (r *userRepositoryImpl) GetUserData(filterBy entities.FilterField, values ...string) (entities.User, error) {
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

func (r *userRepositoryImpl) querySingleField(query string, values []string, user *entities.User) (entities.User, error) {
	if len(values) < 1 {
		return *user, fmt.Errorf("missing query value")
	}
	err := r.db.GetDb().Where(query, values[0]).First(user).Error
	return *user, err
}

func (r *userRepositoryImpl) UpdateUser(in *entities.InsertUserDTO, userId string) (entities.User, error) {
	var user entities.User

	if err := r.db.GetDb().First(&user, "id = ?", userId).Error; err != nil {
		return entities.User{}, err
	}

	if in.Username != "" {
		user.Username = in.Username
	}

	if in.Email != "" {
		user.Email = in.Email
	}

	if in.PhoneNumber != "" {
		user.PhoneNumber = in.PhoneNumber
	}

	if in.FirstName != "" {
		user.FirstName = in.FirstName
	}

	if in.LastName != "" {
		user.LastName = in.LastName
	}

	if in.Password != "" {
		user.Password = in.Password
	}

	if in.ProfilePicture != "" {
		user.ProfileImage = in.ProfilePicture
	}

	user.UpdatedAt = time.Now()

	if err := r.db.GetDb().Save(&user).Error; err != nil {
		return entities.User{}, err
	}

	user.Password = ""

	return user, nil
}
