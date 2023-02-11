package db

import (
	"fmt"
	"imagedisplay/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func COnnectDb() {
	var err error

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	fmt.Println(host)
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s  password=%s ", host, port, user, dbname, password)
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("error connecting to database")

	}
	Db.AutoMigrate(&models.Document{})
}
