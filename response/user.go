package response

import (
	"echo-project/model"
)

// UserResponseHandler - interface for user response
type UserResponseHandlerInterface interface {
	ProcessUserResponse(rid string, d *model.User, code string, err error) interface{}
}

// UserResponse struct
type UserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// UserResponse struct
type UserDetailedResponse struct {
	UserResponse `json:"data"`
	ErrorResponse
}

// Constructor function for UserResponseHandler
func NewUserResponseHandler() UserResponseHandlerInterface {
	return &UserResponse{}
}

// ProcessUserResponse processes the user response
func (u *UserResponse) ProcessUserResponse(rid string, d *model.User, code string, err error) interface{} {
	// r := new(UserDetailedResponse)

	if err == nil {
		return map[string]interface{}{
			"id":    d.ID.Hex(),
			"name":  d.Name,
			"email": d.Email,
		}
	}

	errResp := ProcessErrorResponse(err)
	return errResp
}
