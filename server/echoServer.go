package server

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Bits-Fusion/the_application_backend/config"
	"github.com/Bits-Fusion/the_application_backend/database"
	"github.com/Bits-Fusion/the_application_backend/internal/auth"

	userHandlers "github.com/Bits-Fusion/the_application_backend/features/users/handlers"
	userRepo "github.com/Bits-Fusion/the_application_backend/features/users/repositories"
	userUsecase "github.com/Bits-Fusion/the_application_backend/features/users/usecases"

	taskHandlers "github.com/Bits-Fusion/the_application_backend/features/tasks/handlers"
	taskRepo "github.com/Bits-Fusion/the_application_backend/features/tasks/repositories"
	taskUsecase "github.com/Bits-Fusion/the_application_backend/features/tasks/usecases"

	leadHandlers "github.com/Bits-Fusion/the_application_backend/features/leads/handlers"
	leadRepo "github.com/Bits-Fusion/the_application_backend/features/leads/repositories"
	leadUsecase "github.com/Bits-Fusion/the_application_backend/features/leads/usecases"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

type echoServer struct {
	app  *echo.Echo
	db   database.Database
	conf *config.Config
	auth auth.Authenticator
}

func NewEchoServer(conf *config.Config, db database.Database) *echoServer {
	echoApp := echo.New()
	echoApp.Logger.SetLevel(log.DEBUG)

	return &echoServer{
		app:  echoApp,
		db:   db,
		conf: conf,
		auth: auth.NewJWTAuthenticator(
			conf.TokenConfig.Secret,
			conf.TokenConfig.Iss,
			conf.TokenConfig.Iss,
		),
	}
}

func (s *echoServer) Start() {
	s.app.Use(middleware.Recover())
	s.app.Use(middleware.Logger())

	s.app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PATCH, echo.DELETE, echo.OPTIONS},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
		},
		ExposeHeaders:    []string{echo.HeaderAuthorization},
		AllowCredentials: true,
	}))

	s.app.GET("/v1/health", func(c echo.Context) error {
		return c.String(200, "OK")
	})

	serverUrl := fmt.Sprintf(":%d", s.conf.Server.Port)
	s.initializeRoutes()
	data, err := json.MarshalIndent(s.app.Routes(), "", "  ")
	if err != nil {
		return
	}
	os.WriteFile("routes.json", data, 0644)
	s.app.Logger.Fatal(s.app.Start(serverUrl))
}

func (s *echoServer) initializeRoutes() {
	newUserRepo := userRepo.NewUserPostgresRepository(s.db)
	newUserUsecase := userUsecase.NewUserUsecase(newUserRepo)
	newUserHttp := userHandlers.NewUserHandler(newUserUsecase, s.conf.TokenConfig, s.auth)

	userSign := s.app.Group("/v1/auth")

	userSign.POST("/signup", newUserHttp.SignUp)
	userSign.POST("/login", newUserHttp.SignIn)

	authRouter := s.app.Group("/v1/user")

	authRouter.Use(s.JWTMiddleware)

	authRouter.GET("/", newUserHttp.ListUsers)
	authRouter.GET("/:id", newUserHttp.GetUser)
	authRouter.PATCH("/:id", newUserHttp.UpdateUser)
	authRouter.DELETE("/delete/:id", newUserHttp.DeleteUser)

	newTaskRepo := taskRepo.NewTaskRepository(s.db)
	newTaskUsecase := taskUsecase.NewTaskUsecase(newTaskRepo)
	newTaskHandler := taskHandlers.NewTaskHandler(newTaskUsecase)

	taskRouter := s.app.Group("/v1/task")
	taskRouter.Use(s.JWTMiddleware)

	taskRouter.POST("/", newTaskHandler.CreateTask)
	taskRouter.GET("/", newTaskHandler.ListTasks)
	taskRouter.PATCH("/:taskId", newTaskHandler.UpdateTask)
	taskRouter.DELETE("/:taskId", newTaskHandler.DeleteTask)

	newLeadRepo := leadRepo.NewLeadRepository(s.db)
	newLeadUsecase := leadUsecase.NewLeadUsecase(newLeadRepo)
	newLeadHandler := leadHandlers.NewLeadHandler(newLeadUsecase)

	leadRouter := s.app.Group("/v1/lead")
	leadRouter.Use(s.JWTMiddleware)

	leadRouter.POST("/", newLeadHandler.CreateLead)
	leadRouter.GET("/", newLeadHandler.ListLeads)
	leadRouter.PATCH("/:leadId", newLeadHandler.UpdateLead)
	leadRouter.DELETE("/:leadId", newLeadHandler.DeleteLead)
}
