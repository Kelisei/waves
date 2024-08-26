package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request URL: %s, Method: %s", r.URL.Path, r.Method)
		next.ServeHTTP(w, r)
	})
}

func SetupRoutes() *mux.Router {

	router := mux.NewRouter()
	router.Use(LoggingMiddleware)
	router.HandleFunc("/api/register", RegisterUser).Methods("POST")
	router.HandleFunc("/api/check-register-fields", CheckRegisterFields).Methods("POST")
	return router
}
