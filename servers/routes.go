package servers

import ( // Import the necessary packages
	"echo-project/logger"
	"echo-project/redis"
	"echo-project/servers/controllers" // Import the controllers package
	"echo-project/servers/services"    // Import the services package

	// Import the net/http package
	"github.com/labstack/echo/v4" // Import the echo package
)

// NewRouter is a function that returns a new Echo router
func InitRoutes(e *echo.Echo) { // Define the InitRoutes function that takes an Echo instance as a parameter

	// Create redis connection
	var errorRedis error
	redisClient, errorRedis := redis.NewRedisCacheAdapter() // Create a new RedisCacheAdapter instance
	if errorRedis != nil {
		logger.Error("", "Error connecting to redis", errorRedis)
	}

	// Create monodb connection
	// var errorMongo error

	// mongoClient, errorMongo := mongo.NewMongoDBAdapter() // Create a new MongoDBAdapter instance
	homeService := services.NewHomeService()                                  // Create a new HomeService instance
	homeController := controllers.NewHomeController(homeService, redisClient) // Create a new HomeController instance

	e.GET("/", homeController.Home) // Add a route for the home page
}
