package handlers

import (
	"net/http"
	"strconv"

	"github.com/Bits-Fusion/the_application_backend/features/leads/models"

	"github.com/Bits-Fusion/the_application_backend/features/leads/usecases"
	userModel "github.com/Bits-Fusion/the_application_backend/features/users/models"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
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

func (h *leadHandler) ListLeads(c echo.Context) error {
	limit := c.QueryParam("limit")
	page := c.QueryParam("page")
	orderBy := c.QueryParam("oreder_by")
	assignedTo := c.QueryParam("assigned_to")
	priority := c.QueryParam("priority")
	stage := c.QueryParam("stage")

	var filterOpts models.LeadFilterProps

	filterOpts.OrderBy = orderBy

	if stage != "" {
		filterOpts.Stage = stage
	}

	if priority != "" {
		filterOpts.Priority = priority
	}

	limitInt, _ := strconv.ParseInt(limit, 10, 32)
	pageInt, _ := strconv.ParseInt(page, 10, 32)

	filterOpts.Limit = int32(limitInt)
	filterOpts.Page = int32(pageInt)

	if assignedTo != "" {
		assignedToId, err := uuid.Parse(assignedTo)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
		}

		filterOpts.AssignedTo = assignedToId
	}

	leads, err := h.leadUsecase.ListLeads(filterOpts)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"leads": leads,
	})
}

func (h *leadHandler) UpdateLead(c echo.Context) error {
	var reqBody models.LeadUpdateDTO
	leadId := c.Param("leadId")

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

	lead, err := h.leadUsecase.UpdateLead(&reqBody, leadId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]any{
		"message": "updated successfully",
		"lead":    lead,
	})
}

func (h *leadHandler) DeleteLead(c echo.Context) error {
	id := c.Param("leadId")
	_, err := h.leadUsecase.DeleteLead(userModel.Single, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(200, map[string]string{
		"message": "deleted successfully",
	})
}
