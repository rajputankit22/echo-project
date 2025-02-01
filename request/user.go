package request

import (
	"echo-project/logger"
	"echo-project/model"

	"github.com/labstack/echo/v4"
)

// UserInterface - interface for user
type UserRequestHandlerInterface interface {
	Bind(rid string, c echo.Context, user *model.User, request interface{}) error
}

// UserRequestHandler - struct for user
type UserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

// Constructor function for UserRequestHandler
func NewUserRequestHandler() UserRequestHandlerInterface {
	return &UserRequest{}
}

// Bind function binds the request to the user struct and validates the request
func (u *UserRequest) Bind(rid string, c echo.Context, user *model.User, userRequest interface{}) error {
	var err error
	// Bind request to userStruct
	if err = c.Bind(userRequest); err != nil {
		logger.Error(rid, "Error binding request", err)
		return err
	}
	// Validate user request
	if err = c.Validate(userRequest); err != nil {
		logger.Error(rid, "Error validating request", err)
		return err
	}

	return nil
}
