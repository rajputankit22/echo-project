package controllers

import "github.com/labstack/echo/v4"

// HomeControllerInterface - interface for home controller
type HomeControllerInterface interface {
	Home(c echo.Context) error
}
