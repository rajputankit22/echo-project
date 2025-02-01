package response

import (
	"echo-project/constant"
	"encoding/json"
)

// ErrorResponse struct
type ErrorResponse struct {
	Message string      `json:"message" example:"error message"`
	Code    string      `json:"code" example:"error_code"`
	Error   interface{} `json:"error"`
}

// ProcessErrorResponse processes the error response
func ProcessErrorResponse(err error) interface{} {

	// Marshal the error into JSON
	jsonData, _ := json.MarshalIndent(err, "", "  ")

	// Unmarshal the JSON into a map
	var result map[string]interface{}
	json.Unmarshal(jsonData, &result)

	// Select only the "Fields" field
	fields, _ := result["Fields"].(map[string]interface{})

	return &ErrorResponse{
		Code:    constant.ClientFailureStatusCode,
		Message: constant.StatusText[constant.ClientFailureStatusCode],
		Error:   fields,
	}
}
