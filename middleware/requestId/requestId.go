package requestid

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// RequestIDMiddleware generates a UUID for each request
func RequestIDMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Generate a new UUID
			requestID := uuid.New().String()

			// Set the Request ID in context for later use
			c.Set("request_id", requestID)

			// Add Request ID to Response Headers
			c.Response().Header().Set(echo.HeaderXRequestID, requestID)

			// Call next handler in the chain
			return next(c)
		}
	}
}
