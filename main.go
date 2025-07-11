package main

import (
	"github.com/Bits-Fusion/the_application_backend/config"
	"github.com/Bits-Fusion/the_application_backend/database"
	"github.com/Bits-Fusion/the_application_backend/server"
)

func main() {
	conf := config.GetConfig()
	db := database.NewPostgresDatabase(conf)
	server.NewEchoServer(conf, db).Start()
}
