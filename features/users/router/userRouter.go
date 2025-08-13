package router

import (
	"github.com/Bits-Fusion/the_application_backend/config"
	"github.com/Bits-Fusion/the_application_backend/database"
	"github.com/Bits-Fusion/the_application_backend/features/permissions/entities"
	"github.com/Bits-Fusion/the_application_backend/internal/auth"
	"github.com/labstack/echo/v4"

	permissionRepo "github.com/Bits-Fusion/the_application_backend/features/permissions/repositories"
	userHandlers "github.com/Bits-Fusion/the_application_backend/features/users/handlers"
	userRepo "github.com/Bits-Fusion/the_application_backend/features/users/repositories"
	userUsecase "github.com/Bits-Fusion/the_application_backend/features/users/usecases"
)

type UserRouter interface {
	Init(
		privateRouter *echo.Group,
		publicRouter *echo.Group,
		auth auth.Authenticator,
		tokenConfig *config.TokenConfig,
		permissionMiddleware func(action entities.Action, resource string) echo.MiddlewareFunc,
	)
}

type router struct {
	db database.Database
}

func NewUserRouter(db database.Database) *router {
	return &router{db: db}
}

func (r *router) Init(
	privateRouter *echo.Group,
	publicRouter *echo.Group,
	auth auth.Authenticator,
	tokenConfig *config.TokenConfig,
	permissionMiddleware func(action entities.Action, resource string) echo.MiddlewareFunc,
) {
	newPermissionRepo := permissionRepo.NewPermissionRepository(r.db)

	newUserRepo := userRepo.NewUserPostgresRepository(r.db)
	newUserUsecase := userUsecase.NewUserUsecase(newUserRepo, newPermissionRepo)
	newUserHttp := userHandlers.NewUserHandler(newUserUsecase, tokenConfig, auth)

	publicRouter.POST("/signup", newUserHttp.SignUp)
	publicRouter.POST("/login", newUserHttp.SignIn)

	privateRouter.GET("/", newUserHttp.ListUsers, permissionMiddleware(entities.ActionView, "user"))
	privateRouter.GET("/:id", newUserHttp.GetUser, permissionMiddleware(entities.ActionView, "user"))
	privateRouter.PATCH("/:id", newUserHttp.UpdateUser)
	privateRouter.DELETE("/delete/:id", newUserHttp.DeleteUser, permissionMiddleware(entities.ActionDelete, "user"))
}
