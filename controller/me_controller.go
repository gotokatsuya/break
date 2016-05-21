package controller

import (
	"net/http"

	"github.com/labstack/echo"

	parameter "github.com/gotokatsuya/break/parameter/me"
	service "github.com/gotokatsuya/break/service/me"
)

// LoginMe ...
func LoginMe(ctx echo.Context) error {
	request, err := parameter.NewLoginRequest(ctx)
	if err != nil {
		return err
	}
	response, err := service.LoginMe(ctx, request)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, response)
}

// GetMe ...
func GetMe(ctx echo.Context) error {
	request, err := parameter.NewGetRequest(ctx)
	if err != nil {
		return err
	}
	response, err := service.GetMe(ctx, request)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, response)
}
