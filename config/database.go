package config

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"main.go/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("erp.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Not connect to DB!", err)
	}

	DB = database
	database.AutoMigrate(&models.Company{})

	log.Println("DB connection succesfully!")

}
