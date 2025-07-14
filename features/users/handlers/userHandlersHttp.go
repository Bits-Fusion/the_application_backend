package handlers

import (
	"github.com/Bits-Fusion/the_application_backend/features/users/models"
	"github.com/Bits-Fusion/the_application_backend/features/users/usecases"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type userHandlerImpl struct {
	userUsecase usecases.UserUsecase
}

var validate = validator.New()

func NewUserHandler(usecase usecases.UserUsecase) *userHandlerImpl {
	return &userHandlerImpl{
		userUsecase: usecase,
	}
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
