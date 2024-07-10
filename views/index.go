package views

import (
	"app/templates"
	"context"
	"net/http"

	"github.com/gorilla/sessions"
	"gorm.io/gorm"
)

func LandingPageView(store *sessions.CookieStore, db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// session, _ := store.Get(r, "session-name")
		// userEmail := session.Values["user_email"]
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)

		err := templates.Base("PASS THE GATES", templates.LandingPage()).Render(context.Background(), w)
		if err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			return
		}
	}
}
