package main

import (
	"app/model"
	"app/views"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading the .env: %v", err)
	}

	secretKey := os.Getenv("SECRET_KEY")
	var store = sessions.NewCookieStore([]byte(secretKey))

	dsn := "db.sqlite3"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error al conectar a la base de datos:", err)
		os.Exit(1)
	}
	db.AutoMigrate(&model.User{})
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	r := mux.NewRouter()
	r.HandleFunc("/", views.IndexHandler(store, db)).Methods("GET")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
