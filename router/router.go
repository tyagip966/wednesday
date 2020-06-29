package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"wednesday/container"
)

func InitRouter(e *echo.Echo,controller *container.Container) {
	e.Use(middleware.Logger())
	e.GET("/get_rides",controller.GetCabController().GetRides())
}