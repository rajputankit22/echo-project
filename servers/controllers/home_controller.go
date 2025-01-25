package controllers

import (
	"echo-project/servers/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type homeController struct {
	service services.HomeServiceInterface
}

func NewHomeController(service services.HomeServiceInterface) HomeControllerInterface {
	return &homeController{service: service}
}

func (h *homeController) Home(c echo.Context) error {
	message := h.service.GetWelcomeMessage()
	return c.String(http.StatusOK, message)
}
