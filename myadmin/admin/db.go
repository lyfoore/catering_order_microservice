package admin

import (
	"github.com/theplant/myadmin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var dbParamsString = "user=admin password=123 dbname=admin_dev sslmode=disable host=localhost port=6432"

func ConnectDB() (db *gorm.DB) {
	var err error
	// Create database connection
	db, err = gorm.Open(postgres.Open(dbParamsString))
	if err != nil {
		panic(err)
	}

	// Set db log level
	db.Logger = db.Logger.LogMode(logger.Info)

	// Create data table in the database
	err = db.AutoMigrate(models.Task{})
	if err != nil {
		panic(err)
	}

	return
}
