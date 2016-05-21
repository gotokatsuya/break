package database

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"

	"github.com/gotokatsuya/break/config"
	"github.com/gotokatsuya/break/library/database"
)

// Get ...
func Get(ctx echo.Context) *gorm.DB {
	value := ctx.Get("database")
	if value == nil {
		db := database.NewGorm()
		if config.DebugMode() {
			db = db.Debug()
		}
		ctx.Set("database", db)
		return db
	}
	return value.(*gorm.DB)
}

// Begin ...
func Begin(ctx echo.Context) *gorm.DB {
	txn := Get(ctx).Begin()
	ctx.Set("database", txn)
	return txn
}

// Commit ...
func Commit(ctx echo.Context) {
	txn := Get(ctx).Commit()
	ctx.Set("database", txn)
}

// Rollback ...
func Rollback(ctx echo.Context) {
	txn := Get(ctx).Rollback()
	ctx.Set("database", txn)
}
