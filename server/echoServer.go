package server

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Bits-Fusion/the_application_backend/config"
	"github.com/Bits-Fusion/the_application_backend/database"

	userHandlers "github.com/Bits-Fusion/the_application_backend/features/users/handlers"
	userRepo "github.com/Bits-Fusion/the_application_backend/features/users/repositories"
	userUsecase "github.com/Bits-Fusion/the_application_backend/features/users/usecases"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

type echoServer struct {
	app  *echo.Echo
	db   database.Database
	conf *config.Config
}

func NewEchoServer(conf *config.Config, db database.Database) Server {
	echoApp := echo.New()
	echoApp.Logger.SetLevel(log.DEBUG)

	return &echoServer{
		app:  echoApp,
		db:   db,
		conf: conf,
	}
}

func (s *echoServer) Start() {
	log.Print("starting server")
	s.app.Use(middleware.Recover())
	s.app.Use(middleware.Logger())

	s.app.GET("/v1/health", func(c echo.Context) error {
		return c.String(200, "OK")
	})

	serverUrl := fmt.Sprintf(":%d", s.conf.Server.Port)
	s.initializeUserHandlers()
	data, err := json.MarshalIndent(s.app.Routes(), "", "  ")
	if err != nil {
		return
	}
	os.WriteFile("routes.json", data, 0644)
	s.app.Logger.Fatal(s.app.Start(serverUrl))
}

func (s *echoServer) initializeUserHandlers() {
	newUserRepo := userRepo.NewUserPostgresRepository(s.db)

	newUserUsecase := userUsecase.NewUserUsecase(
		newUserRepo,
	)
	newUserHttp := userHandlers.NewUserHandler(newUserUsecase)
	userRouters := s.app.Group("/v1/auth")
	userRouters.POST("", newUserHttp.SignUp)
}
