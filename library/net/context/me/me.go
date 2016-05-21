package me

import (
	"github.com/labstack/echo"

	"github.com/gotokatsuya/break/model/user"
)

func Get(ctx echo.Context) *user.Entity {
	value := ctx.Get("me")
	if value == nil {
		return nil
	}
	return value.(*user.Entity)
}

func Set(ctx echo.Context, me *user.Entity) {
	ctx.Set("me", me)
}
