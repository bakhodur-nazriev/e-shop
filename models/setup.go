package models

import "github.com/jinzhu/gorm"

var DB *gorm.DB

func ConnectToDatabase() {
	database, err := gorm.Open("postgres", "user=bakhodur password=1996 dbname=e-shop sslmode=disable")
	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Product{}, &User{}, &Categories{})

	DB = database
}
