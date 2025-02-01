package controllers

import (
	"echo-project/logger"
	"echo-project/model"
	"echo-project/mongodb"
	"echo-project/redis"
	"echo-project/request"
	"echo-project/servers/services"
	"fmt"
	"net/http"

	// "github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
)

// struct homeController
type homeController struct {
	service     services.HomeServiceInterface
	redisClient redis.CacheInterface
	mongoClient mongodb.MongoDBInterface
	userRequest request.UserRequestHandlerInterface
}

// Constructor function for homeController
func NewHomeController(service services.HomeServiceInterface, redisClient redis.CacheInterface, mongoClient mongodb.MongoDBInterface) HomeControllerInterface {
	return &homeController{
		service:     service,
		redisClient: redisClient,
		mongoClient: mongoClient,
		userRequest: request.NewUserRequestHandler(),
	}
}

// Home function returns a welcome message
func (h *homeController) Home(c echo.Context) error {
	message := h.service.GetWelcomeMessage()
	if err := h.redisClient.Set("aa", []byte("welcome_message"), 1800); err != nil {
		logger.Error("", "Error setting value in cache", err)
		return c.String(http.StatusInternalServerError, err.Error())
	}

	user := new(model.User)
	userRequest := new(request.UserRequest)

	if errValidator := h.userRequest.Bind("1", c, user, userRequest); errValidator != nil {
		logger.Error("", "Error binding and validating request", errValidator)
		return c.String(http.StatusInternalServerError, errValidator.Error())
	}

	fmt.Println("Request:-------", h)

	logger.Info("", "Error setting value in cache------------------")
	return c.String(http.StatusOK, message)
}
