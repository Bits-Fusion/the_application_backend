package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Bits-Fusion/the_application_backend/config"
	"github.com/Bits-Fusion/the_application_backend/features/users/entities"
	"github.com/Bits-Fusion/the_application_backend/features/users/models"
	"github.com/Bits-Fusion/the_application_backend/features/users/usecases"
	"github.com/Bits-Fusion/the_application_backend/internal/auth"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"

	"github.com/go-playground/validator/v10"
)

type userHandlerImpl struct {
	userUsecase usecases.UserUsecase
	config      *config.TokenConfig
	auth        auth.Authenticator
}

type CreateUserTokenPayload struct {
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,min=3,max=72"`
}

var validate = validator.New()

func NewUserHandler(usecase usecases.UserUsecase, config *config.TokenConfig, auth auth.Authenticator) *userHandlerImpl {
	return &userHandlerImpl{
		userUsecase: usecase,
		config:      config,
		auth:        auth,
	}
}

func (h *userHandlerImpl) SignIn(c echo.Context) error {
	var reqBody CreateUserTokenPayload

	if err := c.Bind(&reqBody); err != nil {
		log.Errorf("Error binding request body: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid request format",
		})
	}

	if err := validate.Struct(reqBody); err != nil {
		errs := make(map[string]string)

		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			for _, fieldErr := range validationErrs {
				errs[fieldErr.Field()] = fieldErr.Error()
			}
		} else {
			errs["general"] = err.Error()
		}

		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "Validation failed",
			"errors":  errs,
		})
	}

	user, err := h.userUsecase.GetUserData(entities.FilterByEmail, reqBody.Email)

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "User not found"})
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(reqBody.Password)) != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid credentials"})
	}

	claims := jwt.MapClaims{
		"sub": user.Id,
		"exp": time.Now().Add(h.config.Exp).Unix(),
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
		"iss": h.config.Iss,
		"aud": h.config.Iss,
	}

	token, err := h.auth.GenerateToken(claims)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Can not create auth token"})
	}

	return c.JSON(http.StatusCreated, map[string]string{"token": token})

}

func (h *userHandlerImpl) SignUp(c echo.Context) error {
	var reqBody models.UserModel

	if err := c.Bind(&reqBody); err != nil {
		log.Errorf("Error binding request body: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid request format",
		})
	}

	if err := validate.Struct(reqBody); err != nil {
		errs := make(map[string]string)

		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			for _, fieldErr := range validationErrs {
				errs[fieldErr.Field()] = fieldErr.Error()
			}
		} else {
			errs["general"] = err.Error()
		}

		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "Validation failed",
			"errors":  errs,
		})
	}

	if err := h.userUsecase.CreateUser(&reqBody); err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Internal server error",
		})
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"message": "User created successfully",
	})
}

func (h *userHandlerImpl) ListUsers(e echo.Context) error {
	limit := e.QueryParams().Get("limit")
	page := e.QueryParams().Get("page")
	orderBy := e.QueryParams().Get("oreder_by")

	var param entities.FilterParams
	param.OrderBy = orderBy

	limitInt, _ := strconv.ParseInt(limit, 10, 32)
	pageInt, _ := strconv.ParseInt(page, 10, 32)

	param.Limit = int32(limitInt)
	param.Page = int32(pageInt)

	users, err := h.userUsecase.ListUser(param)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]error{
			"error": err,
		})
	}

	return e.JSON(http.StatusOK, map[string][]entities.User{
		"users": users,
	})
}
