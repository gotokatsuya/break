package database

import (
	log "github.com/Sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"github.com/gotokatsuya/break/config"
)

func newGorm() *gorm.DB {
	// Open
	db, err := gorm.Open(config.GetDB().Driver, config.GetDB().Source)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Panic("Database:Open")
	}
	db.DB().SetMaxIdleConns(10)
	return db
}

var gormDB *gorm.DB

// NewGorm ...
func NewGorm() *gorm.DB {
	if gormDB == nil {
		gormDB = newGorm()
	}
	return gormDB
}
