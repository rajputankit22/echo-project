package main

import (
	"echo-project/logger"
	requestid "echo-project/middleware/requestId"
	"echo-project/middleware/validator"
	"echo-project/servers"

	"echo-project/config"

	"github.com/labstack/echo/v4"
)

func init() {
	// logger.Trace("Starting the application")
}

func main() {
	e := echo.New() // Create a new Echo instance

	e.Validator = validator.NewValidatorAdapter() // Create a new Echo instance

	// Initialize the logger
	defer logger.CloseLogger() // Close the logger when done

	// Middleware
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	// Add Custom Request ID Middleware
	e.Use(requestid.RequestIDMiddleware())

	logger.Trace("Starting the application")
	servers.InitRoutes(e)                             // Initialize the routes
	e.Logger.Fatal(e.Start(config.Config().HTTPPort)) // Start the server
}
