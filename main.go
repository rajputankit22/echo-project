package main

import (
	"echo-project/logger"
	"echo-project/servers"

	"echo-project/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	logger.Info("123", "Starting the application")
}

func main() {
	e := echo.New() // Create a new Echo instance

	// Initialize the logger
	defer logger.CloseLogger() // Close the logger when done

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	servers.InitRoutes(e)                             // Initialize the routes
	e.Logger.Fatal(e.Start(config.Config().HTTPPort)) // Start the server
}
