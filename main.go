package main

import (
	"go-backend/config"
	"go-backend/handlers"
	"go-backend/repositories"
	"go-backend/routes"
	"go-backend/services"
)

func main() {
	config.ConnectDatabase()

	roleRepo := repositories.NewRoleRepository(config.DB)
	roleService := services.NewRoleService(roleRepo)
	roleHandler := handlers.NewRoleHandler(roleService)

	r := routes.SetupRouter(roleHandler)
	r.Run(":8080")
}
