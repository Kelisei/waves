package api

import (
	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/home", HomeHandler).Methods("GET")
	router.HandleFunc("/api/login", LoginHandler).Methods("POST")
	router.HandleFunc("/api/logout", LogoutHandler).Methods("POST")
	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
	return router
}
