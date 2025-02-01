package controllers

import (
	"echo-project/constant"
	"echo-project/logger"
	"echo-project/model"
	"echo-project/mongodb"
	"echo-project/redis"
	"echo-project/request"
	"echo-project/response"
	"echo-project/servers/services"
	"net/http"

	// "github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
)

// struct homeController
type homeController struct {
	service      services.HomeServiceInterface
	redisClient  redis.CacheInterface
	mongoClient  mongodb.MongoDBInterface
	userRequest  request.UserRequestHandlerInterface
	userResponse response.UserResponseHandlerInterface
}

// Constructor function for homeController
func NewHomeController(service services.HomeServiceInterface, redisClient redis.CacheInterface, mongoClient mongodb.MongoDBInterface) HomeControllerInterface {
	return &homeController{
		service:      service,
		redisClient:  redisClient,
		mongoClient:  mongoClient,
		userRequest:  request.NewUserRequestHandler(),
		userResponse: response.NewUserResponseHandler(),
	}
}

// Home function returns a welcome message
func (h *homeController) Home(c echo.Context) error {

	var err error
	// Retrieve Request ID from context
	requestID := c.Get("request_id").(string)
	logger.Info(requestID, "Request received")

	message := h.service.GetWelcomeMessage()
	if err := h.redisClient.Set("aa", []byte("welcome_message"), 1800); err != nil {
		logger.Error(requestID, "Error setting value in cache", err)
		return c.String(http.StatusInternalServerError, err.Error())
	}

	user := new(model.User)
	userRequest := new(request.UserRequest)

	if err = h.userRequest.Bind(requestID, c, user, userRequest); err != nil {
		logger.Error(requestID, "Error binding and validating request", err)
	}

	if err != nil {
		errResp := h.userResponse.ProcessUserResponse(requestID, nil, "", err).(*response.ErrorResponse)
		return c.JSON(constant.HttpStatusCodes[errResp.Code], errResp)
	}

	return c.String(http.StatusOK, message)
}
