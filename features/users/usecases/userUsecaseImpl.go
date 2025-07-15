package usecases

import (
	"errors"
	"regexp"
	"strings"

	"github.com/Bits-Fusion/the_application_backend/features/users/entities"
	"github.com/Bits-Fusion/the_application_backend/features/users/models"
	"github.com/Bits-Fusion/the_application_backend/features/users/repositories"
	"golang.org/x/crypto/bcrypt"
)

type userUsecaseImpl struct {
	UserRepository repositories.UserRepository
}

func NewUserUsecase(
	userRepository repositories.UserRepository,
) *userUsecaseImpl {
	return &userUsecaseImpl{
		UserRepository: userRepository,
	}
}

func StandardizePhoneNumber(phoneNumber string) (string, error) {
	if strings.TrimSpace(phoneNumber) == "" {
		return "", errors.New("invalid phone number")
	}

	cleaned := strings.ReplaceAll(phoneNumber, " ", "")
	cleaned = strings.ReplaceAll(cleaned, "-", "")

	lengthCheck := func(num string, expected int) bool {
		return len(num) == expected
	}

	if matched, _ := regexp.MatchString(`^(09|07)`, cleaned); matched && lengthCheck(cleaned, 10) {
		return "+251" + cleaned[1:], nil
	}

	if (strings.HasPrefix(cleaned, "+251") || strings.HasPrefix(cleaned, "251")) &&
		lengthCheck(strings.TrimPrefix(cleaned, "+"), 12) {
		if strings.HasPrefix(cleaned, "251") {
			return "+" + cleaned, nil
		}
		return cleaned, nil
	}

	if matched, _ := regexp.MatchString(`^[97]`, cleaned); matched && lengthCheck(cleaned, 9) {
		return "+251" + cleaned, nil
	}

	return "", errors.New("invalid phone number")
}

func (r *userUsecaseImpl) CreateUser(in *models.UserModel) error {

	phoneNumber, err := StandardizePhoneNumber(in.PhoneNumber)

	if err != nil {
		return err
	}

	if _, err := r.UserRepository.GetUserData(entities.FilterByAll, in.Username, in.Email, phoneNumber); err == nil {
		return errors.New("user with this credential exist")
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(in.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	newDto := &entities.InsertUserDTO{
		Username:    in.Username,
		Email:       in.Email,
		PhoneNumber: phoneNumber,
		Password:    string(encryptedPassword),
		FirstName:   in.FirstName,
		LastName:    in.LastName,
		Role:        "admin",
	}

	if err := r.UserRepository.InsertUserData(newDto); err != nil {
		return err
	}

	return nil
}

func (u *userUsecaseImpl) ListUser(params entities.FilterParams) ([]entities.User, error) {
	return u.UserRepository.ListUsers(params)
}

func (u *userUsecaseImpl) GetUserData(filterBy entities.FilterField, values ...string) (entities.User, error) {
	return u.UserRepository.GetUserData(filterBy, values...)
}
