package main

import (
	"github.com/henvo/golang-ddd-starter/internal/infrastructure/postgres"
	"github.com/henvo/golang-ddd-starter/internal/interface/echo"
)

func main() {

	// Create the echo web server instance.
	web := echo.NewWebServer()

	// User service
	userService := postgres.NewUserService()
	userController := echo.NewUserController(userService)
	web.Add(userController)

	// Start the server.
	web.Start(":8080")
}
