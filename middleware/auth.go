package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/labstack/echo"

	"github.com/gotokatsuya/break/library/net/context/me"
	"github.com/gotokatsuya/break/model/user"
)

var (
	errNoAccessToken      = errors.New("Access token is not found.")
	errInvalidAccessToken = errors.New("Access token is invalid.")
)

// AuthRequired API用のAccessTokenが正しいか確認する
func AuthRequired() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		auth := ctx.Request().Header().Get("Authorization")
		if err := validateAccessToken(ctx, parseAccessToken(auth)); err != nil {
			ctx.JSON(http.StatusUnauthorized, err)
			return err
		}
		return nil
	}
}

func parseAccessToken(auth string) string {
	if !strings.HasPrefix(auth, "AccessToken ") {
		return ""
	}
	return strings.TrimPrefix(auth, "AccessToken ")
}

func validateAccessToken(ctx echo.Context, accessToken string) error {
	if len(accessToken) <= 0 {
		return errNoAccessToken
	}
	userRepository := user.NewRepository(ctx)
	user, err := userRepository.GetByToken(accessToken)
	if user == nil || err != nil {
		return errInvalidAccessToken
	}
	me.Set(ctx, user)
	return nil
}
