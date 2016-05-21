package main

import (
	"github.com/labstack/echo/engine/fasthttp"

	"github.com/gotokatsuya/break/config"
	"github.com/gotokatsuya/break/library/database"
	"github.com/gotokatsuya/break/library/database/scheme"
	"github.com/gotokatsuya/break/route"
)

func init() {
	config.Load()
	db := database.NewGorm()
	scheme.Sync(db)
}

func main() {
	handler := route.Handler()
	switch {
	case config.DebugMode():
		handler.SetDebug(true)
	}
	handler.Run(fasthttp.New(":8888"))
}
