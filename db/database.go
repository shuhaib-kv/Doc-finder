package db

import (
	"fmt"
	"imagedisplay/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func COnnectDb() {
	var err error
	dsn := "host=localhost port=5432 user=postgres password=soib  dbname=docs   "
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("error connecting to database")

	}
	Db.AutoMigrate(&models.Document{})
}
