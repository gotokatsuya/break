package controller

import (
	"net/http"

	"github.com/labstack/echo"

	parameter "github.com/gotokatsuya/break/parameter/spot"
	service "github.com/gotokatsuya/break/service/spot"
)

// CreateSpot ...
func CreateSpot(ctx echo.Context) error {
	request, err := parameter.NewCreateRequest(ctx)
	if err != nil {
		return err
	}
	if err := service.CreateSpot(ctx, request); err != nil {
		return err
	}
	return ctx.JSON(http.StatusCreated, map[string]interface{}{})
}

// GetSpot ...
func GetSpot(ctx echo.Context) error {
	request, err := parameter.NewGetRequest(ctx)
	if err != nil {
		return err
	}
	res, err := service.GetSpot(ctx, request)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, res)
}
