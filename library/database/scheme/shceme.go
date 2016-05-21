package scheme

import (
	"github.com/jinzhu/gorm"

	spotmodel "github.com/gotokatsuya/break/model/spot"
	tourmodel "github.com/gotokatsuya/break/model/tour"
	tourspotmodel "github.com/gotokatsuya/break/model/tour/tourspot"
	usermodel "github.com/gotokatsuya/break/model/user"
)

func createTables(db *gorm.DB, models []interface{}) {
	for _, model := range models {
		db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4").CreateTable(model)
	}
}

func autoMigrateTables(db *gorm.DB, models []interface{}) {
	for _, model := range models {
		db.AutoMigrate(model)
	}
}

func Sync(db *gorm.DB) {
	models := []interface{}{
		new(usermodel.Entity),
		new(spotmodel.Entity),
		new(tourmodel.Entity),
		new(tourspotmodel.Entity),
	}
	createTables(db, models)
	autoMigrateTables(db, models)
}
