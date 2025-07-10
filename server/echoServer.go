package server

import (
 "fmt"

 "github.com/Bits-Fusion/the_application_backend/config"
 "github.com/Bits-Fusion/the_application_backend/database"
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
  s.app.Use(middleware.Recover())
  s.app.Use(middleware.Logger())

  // Health check adding
  s.app.GET("v1/health", func(c echo.Context) error {
    return c.String(200, "OK")
  })
  
  serverUrl := fmt.Sprintf(":%d", s.conf.Server.Port)
  s.app.Logger.Fatal(s.app.Start(serverUrl))
}
