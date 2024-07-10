package views

import (
	"app/templates"
	"context"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"gorm.io/gorm"
)

func IndexHandler(store *sessions.CookieStore, db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// session, _ := store.Get(r, "session-name")
		// userEmail := session.Values["user_email"]
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		log.Fatal(templates.Hello("Roman Riquelme").Render(context.Background(), w))

	}
}
