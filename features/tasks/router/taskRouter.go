package router

import (
	"github.com/Bits-Fusion/the_application_backend/database"
	"github.com/Bits-Fusion/the_application_backend/features/permissions/entities"
	"github.com/labstack/echo/v4"

	taskHandlers "github.com/Bits-Fusion/the_application_backend/features/tasks/handlers"
	taskRepo "github.com/Bits-Fusion/the_application_backend/features/tasks/repositories"
	taskUsecase "github.com/Bits-Fusion/the_application_backend/features/tasks/usecases"
)

type TaskRouter interface {
	Init(
		leadRouterGroupGroup *echo.Group,
		permissionMiddleware func(action entities.Action, resource string) echo.MiddlewareFunc,
	)
}

type router struct {
	db database.Database
}

func NewTaskRouter(db database.Database) *router {
	return &router{db: db}
}

func (r *router) Init(
	taskRouterGroup *echo.Group,
	permissionMiddleware func(action entities.Action, resource string) echo.MiddlewareFunc,
) {

	newTaskRepo := taskRepo.NewTaskRepository(r.db)
	newTaskUsecase := taskUsecase.NewTaskUsecase(newTaskRepo)
	newTaskHandler := taskHandlers.NewTaskHandler(newTaskUsecase)

	taskRouterGroup.POST("/", newTaskHandler.CreateTask, permissionMiddleware(entities.ActionCreate, "task"))
	taskRouterGroup.GET("/", newTaskHandler.ListTasks, permissionMiddleware(entities.ActionView, "task"))
	taskRouterGroup.PATCH("/:taskId", newTaskHandler.UpdateTask, permissionMiddleware(entities.ActionUpdate, "task"))
	taskRouterGroup.DELETE("/:taskId", newTaskHandler.DeleteTask, permissionMiddleware(entities.ActionDelete, "task"))
}
