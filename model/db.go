package model

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DB es la instancia global de la base de datos
var DB *gorm.DB

// InitDB inicializa la conexión a la base de datos
func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	// Migrar los modelos (puedes añadir otros modelos aquí)
	DB.AutoMigrate(&User{})
}
