package route

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/gotokatsuya/break/controller"
	breakmiddleware "github.com/gotokatsuya/break/middleware"
)

// Handler ...
func Handler() *echo.Echo {
	router := echo.New()
	router.Use(middleware.Recover())
	router.Use(middleware.Logger())
	router.Use(middleware.Gzip())
	router.Use(middleware.Secure())
	apiGroup := router.Group("/api")
	{
		apiGroup.POST("/me/login", controller.LoginMe)
		authorizedGroup := apiGroup.Group("/")
		authorizedGroup.Use(echo.WrapMiddleware(breakmiddleware.AuthRequired()))
		{
			authorizedGroup.GET("me", controller.GetMe)
			authorizedGroup.GET("spot", controller.GetSpot)
			authorizedGroup.POST("spot", controller.CreateSpot)
			authorizedGroup.GET("tour", controller.GetTour)
			authorizedGroup.POST("tour", controller.CreateTour)
		}
	}
	return router
}
