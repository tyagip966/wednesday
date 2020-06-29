package controller

import (
	"github.com/labstack/echo"
	"net/http"
	"wednesday/models"
)

type CabController struct {
	cabUseCase   models.CabUseCase
}

func NewCabController(cabUseCase models.CabUseCase) *CabController {
	return &CabController{cabUseCase: cabUseCase}
}

func (o *CabController) GetRides() echo.HandlerFunc{
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	}
}
