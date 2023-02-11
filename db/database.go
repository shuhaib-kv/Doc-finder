package db

import (
	"fmt"
	"imagedisplay/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBS *gorm.DB

func COnnectDb() {
	var err error
	DBS, err = gorm.Open(postgres.Open("host=localhost port=5432 user=postgres password=soib dbname=interview"), &gorm.Config{})
	if err != nil {
		fmt.Println("Datatbase connection faild")
	}
	DBS.AutoMigrate(
		&models.Document{},
	)
	DBS.AutoMigrate(&models.Document{})
}
