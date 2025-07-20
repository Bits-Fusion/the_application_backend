package handlers

import (
	"net/http"
	"strconv"

	"github.com/Bits-Fusion/the_application_backend/features/tasks/entities"
	"github.com/Bits-Fusion/the_application_backend/features/tasks/models"
	"github.com/Bits-Fusion/the_application_backend/features/tasks/usecases"
	userModel "github.com/Bits-Fusion/the_application_backend/features/users/models"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type taskHandlerImpl struct {
	taskUsecase usecases.TaskUsecase
}

var validate = validator.New()

func NewTaskHandler(taskUsecase usecases.TaskUsecase) *taskHandlerImpl {
	return &taskHandlerImpl{
		taskUsecase: taskUsecase,
	}
}

func (h *taskHandlerImpl) CreateTask(c echo.Context) error {
	var reqBody models.TaskModel

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

	if err := h.taskUsecase.CreateTask(&reqBody); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]any{
		"message": "created successfully",
	})
}

func (h *taskHandlerImpl) ListTasks(c echo.Context) error {
	limit := c.QueryParam("limit")
	page := c.QueryParam("page")
	orderBy := c.QueryParam("oreder_by")
	assignedTo := c.QueryParam("assigned_to")
	priority := c.QueryParam("priority")
	status := c.QueryParam("status")

	var filterOpts models.TaskFilterProps

	filterOpts.OrderBy = orderBy

	if status != "" {
		filterOpts.Status = models.StatusFiterOpt(status)
	}

	if priority != "" {
		filterOpts.Priority = models.PriorityFilterOpt(priority)
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

	tasks, err := h.taskUsecase.ListTask(filterOpts)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(200, map[string][]entities.Task{
		"tasks": tasks,
	})
}

func (h *taskHandlerImpl) UpdateTask(c echo.Context) error {
	taskId := c.Param("taskId")

	var reqBody models.TaskModelUpdate

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

	updatedTask, err := h.taskUsecase.UpdateTask(&reqBody, taskId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"task": updatedTask,
	})
}

func (h *taskHandlerImpl) DeleteTask(c echo.Context) error {
	taskId := c.Param("taskId")

	if _, err := h.taskUsecase.DeleteTask(userModel.Single, taskId); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "deleted successfully",
	})
}
