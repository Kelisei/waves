package main

import (
	"app/model"
	"app/templates"
	"context"
	"fmt"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	dsn := "db.sqlite3"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error al conectar a la base de datos:", err)
		os.Exit(1)
	}
	fmt.Println("This is a test")
	db.AutoMigrate(&model.User{})

	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	y := model.NewUser("Goku", "goku@gmail.com", "Rkt123@", "Argentina", nil, nil, nil)
	db.Create(&y)
	x := templates.Hello(y.Name)
	fmt.Println(x.Render(context.Background(), os.Stdout))
}
