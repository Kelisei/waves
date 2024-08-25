package main

import (
	"fmt"
	"net/http"
	"waves/api"
	"waves/model"

	"github.com/rs/cors"
)

func main() {

	model.InitDB()
	handler := cors.Default().Handler(api.SetupRoutes())

	fmt.Println("Starting server at http://localhost:8081/")
	if err := http.ListenAndServe(":8081", handler); err != nil {
		fmt.Println(err)
	}
}
