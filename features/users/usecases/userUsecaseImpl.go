package usecases

import (
	"errors"
	"os"
	"regexp"
	"strings"

	permissionEntity "github.com/Bits-Fusion/the_application_backend/features/permissions/entities"
	permissionRepo "github.com/Bits-Fusion/the_application_backend/features/permissions/repositories"

	"github.com/Bits-Fusion/the_application_backend/features/users/entities"
	"github.com/Bits-Fusion/the_application_backend/features/users/models"
	"github.com/Bits-Fusion/the_application_backend/features/users/repositories"
	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

type userUsecaseImpl struct {
	UserRepository       repositories.UserRepository
	PermissionRepository permissionRepo.PermissionRepository
}

func NewUserUsecase(
	userRepository repositories.UserRepository,
	permissionRepository permissionRepo.PermissionRepository,
) *userUsecaseImpl {
	return &userUsecaseImpl{
		UserRepository:       userRepository,
		PermissionRepository: permissionRepository,
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
		Id:          uuid.New(),
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

func (u *userUsecaseImpl) UpdateUser(in *models.UserUpdateModel, userId string) (entities.User, error) {
	var user entities.InsertUserDTO

	if in.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword(
			[]byte(in.Password),
			bcrypt.DefaultCost,
		)
		if err != nil {
			return entities.User{}, err

		}
		user.Password = string(hashedPassword)
	}

	if in.Username != "" {
		_, err := u.UserRepository.GetUserData(entities.FilterByUsername, in.Username)

		if err == nil {
			return entities.User{}, errors.New("user with this username already exist")
		}

		user.Username = in.Username
	}

	if in.Permission != nil {
		var createdPermissions []permissionEntity.Permission

		allowedActions := make(map[string]struct{})
		allowedActions["view"] = struct{}{}
		allowedActions["create"] = struct{}{}
		allowedActions["update"] = struct{}{}
		allowedActions["delete"] = struct{}{}

		allowedResources := make(map[string]struct{})

		allowedResources["task"] = struct{}{}
		allowedResources["lead"] = struct{}{}
		allowedResources["user"] = struct{}{}

		for _, perm := range in.Permission {
			tot := strings.Split(perm, "_")

			if len(tot) != 2 {
				log.Error("invalid permission type ")
				continue
			}
			if _, ok := allowedActions[tot[1]]; !ok {
				log.Error("invalid permission action")
				continue
			}
			if _, ok := allowedResources[tot[0]]; !ok {
				log.Error("invalid permission resource")
				continue
			}

			permission, err := u.PermissionRepository.CreatePermission(tot[1], tot[0])

			if err != nil {
				log.Error(err)
			}

			createdPermissions = append(createdPermissions, *permission)
		}
		user.Permission = createdPermissions
	}

	if in.Email != "" {
		_, err := u.UserRepository.GetUserData(entities.FilterByEmail, in.Email)

		if err == nil {
			return entities.User{}, errors.New("user with this email already exist")
		}

		user.Email = in.Email
	}

	if in.PhoneNumber != "" {
		standardPhone, err := StandardizePhoneNumber(in.PhoneNumber)

		if err != nil {
			return entities.User{}, errors.New("invalid phone number")
		}

		_, err = u.UserRepository.GetUserData(entities.FilterByPhoneNumber, standardPhone)

		if err == nil {
			return entities.User{}, errors.New("user with this phone number exist")
		}

		user.PhoneNumber = standardPhone
	}

	if in.LastName != "" {
		user.LastName = in.LastName
	}
	if in.FirstName != "" {
		user.FirstName = in.FirstName
	}

	if in.ProfilePicture != "" {
		privUser, err := u.GetUserData(entities.FilterByID, userId)
		if err != nil {
			return entities.User{}, err
		}
		if privUser.ProfileImage != "" {
			if err := os.Remove(privUser.ProfileImage); err != nil {
				log.Error(err)
			}
		}
		user.ProfilePicture = in.ProfilePicture
	}

	updatedUser, err := u.UserRepository.UpdateUser(&user, userId)

	if err != nil {
		return entities.User{}, err
	}

	return updatedUser, nil
}

func (u *userUsecaseImpl) ListUser(params entities.FilterParams) ([]entities.User, error) {
	return u.UserRepository.ListUsers(params)
}

func (u *userUsecaseImpl) GetUserData(filterBy entities.FilterField, values ...string) (entities.User, error) {
	return u.UserRepository.GetUserData(filterBy, values...)
}

func (u *userUsecaseImpl) DeleteUser(deletionMode models.DeleteMode, userId ...string) (bool, error) {
	return u.UserRepository.DeleteUser(deletionMode, userId...)
}
