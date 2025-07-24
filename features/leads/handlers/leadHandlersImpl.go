package handlers

import (
	"net/http"

	"github.com/Bits-Fusion/the_application_backend/features/leads/models"
	"github.com/Bits-Fusion/the_application_backend/features/leads/usecases"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type leadHandler struct {
	leadUsecase usecases.LeadUsecase
}

func NewLeadHandler(leadUsecase usecases.LeadUsecase) *leadHandler {
	return &leadHandler{
		leadUsecase: leadUsecase,
	}
}

var validate = validator.New()

func (h *leadHandler) CreateLead(c echo.Context) error {
	var reqBody models.LeadInsertDTO

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

	if err := h.leadUsecase.CreateLead(&reqBody); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]any{
		"message": "created successfully",
	})
}
