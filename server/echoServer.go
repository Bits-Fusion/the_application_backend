package server

import (
	"fmt"

	"github.com/Bits-Fusion/the_application_backend/config"
	"github.com/Bits-Fusion/the_application_backend/database"
	"github.com/Bits-Fusion/the_application_backend/internal/auth"

	leadRoute "github.com/Bits-Fusion/the_application_backend/features/leads/router"
	taskRoute "github.com/Bits-Fusion/the_application_backend/features/tasks/router"
	userRoute "github.com/Bits-Fusion/the_application_backend/features/users/router"

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
	s.app.Logger.Fatal(s.app.Start(serverUrl))
}

func (s *echoServer) initializeRoutes() {
	// user
	publicUserRoute := s.app.Group("/v1/auth")
	privateUserRoute := s.app.Group("/v1/user")

	privateUserRoute.Use(s.JWTMiddleware)
	newUserRoute := userRoute.NewUserRouter(s.db)
	newUserRoute.Init(
		privateUserRoute,
		publicUserRoute,
		s.auth,
		s.conf.TokenConfig,
		s.RequirePermission,
	)

	// task
	taskGroup := s.app.Group("/v1/task")
	taskGroup.Use(s.JWTMiddleware)

	taskRouter := taskRoute.NewTaskRouter(s.db)
	taskRouter.Init(taskGroup, s.RequirePermission)

	// lead
	leadGroup := s.app.Group("/v1/lead")
	leadGroup.Use(s.JWTMiddleware)

	leadRouter := leadRoute.NewLeadRouter(s.db)
	leadRouter.Init(leadGroup, s.RequirePermission)

}
