package model

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	err = DB.AutoMigrate(
		&User{},
		&Artist{},
		&Genre{},
		&Audio{},
		&Collection{},
	)
	if err != nil {
		log.Fatal("failed to migrate database: ", err)
	}

	log.Println("Database connection and migration successful")
}
