package controllers

import "github.com/labstack/echo/v4"

type HomeControllerInterface interface {
	Home(c echo.Context) error
}
