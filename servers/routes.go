package servers

import ( // Import the necessary packages
	"echo-project/servers/controllers" // Import the controllers package
	"echo-project/servers/services"    // Import the services package

	// Import the net/http package
	"github.com/labstack/echo/v4" // Import the echo package
)

// NewRouter is a function that returns a new Echo router
func InitRoutes(e *echo.Echo) { // Define the InitRoutes function that takes an Echo instance as a parameter
	homeService := services.NewHomeService()                     // Create a new HomeService instance
	homeController := controllers.NewHomeController(homeService) // Create a new HomeController instance

	e.GET("/", homeController.Home) // Add a route for the home page
}
