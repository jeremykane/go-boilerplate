package api

import (
	"github.com/jeremykane/go-boilerplate/internal/config"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func RegisterRoutes(echoServer *echo.Echo, handler *Handler, cfg *config.Config) {
	echoServer.GET("/swagger/*", echoSwagger.WrapHandler)
	echoServer.GET("/facility-types", handler.FacilityTypeHandler.GetAll)
}
