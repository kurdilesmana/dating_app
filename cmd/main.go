package main

import (
	"github.com/kurdilesmana/dating_app/config"
	"github.com/kurdilesmana/dating_app/internal/delivery/http/router"
	"github.com/kurdilesmana/dating_app/internal/infrastructure/database"
)

func main() {
	config.InitConfig()
	database.InitDatabase()
	server := router.InitRouter()

	// Start the server
	server.Start(":8080")
}
