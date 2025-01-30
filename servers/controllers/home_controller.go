package controllers

import (
	"echo-project/logger"
	"echo-project/mongodb"
	"echo-project/redis"
	"echo-project/servers/services"
	"net/http"

	// "github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
)

// struct homeController
type homeController struct {
	service     services.HomeServiceInterface
	redisClient redis.CacheInterface
	mongoClient mongodb.MongoDBInterface
}

// Constructor function for homeController
func NewHomeController(service services.HomeServiceInterface, redisClient redis.CacheInterface, mongoClient mongodb.MongoDBInterface) HomeControllerInterface {
	return &homeController{
		service:     service,
		redisClient: redisClient,
		mongoClient: mongoClient,
	}
}

// Home function returns a welcome message
func (h *homeController) Home(c echo.Context) error {
	message := h.service.GetWelcomeMessage()
	if err := h.redisClient.Set("aa", []byte("welcome_message"), 1800); err != nil {
		logger.Error("", "Error setting value in cache", err)
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, message)
}
