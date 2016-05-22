package controller

import (
	"net/http"

	"github.com/labstack/echo"

	parameter "github.com/gotokatsuya/break/parameter/tour"
	service "github.com/gotokatsuya/break/service/tour"
)

// CreateTour ...
func CreateTour(ctx echo.Context) error {
	request, err := parameter.NewCreateRequest(ctx)
	if err != nil {
		return err
	}
	if err := service.CreateTour(ctx, request); err != nil {
		return err
	}
	return ctx.JSON(http.StatusCreated, map[string]interface{}{})
}

// GetTour ...
func GetTour(ctx echo.Context) error {
	request, err := parameter.NewGetRequest(ctx)
	if err != nil {
		return err
	}
	res, err := service.GetTour(ctx, request)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, res)
}
