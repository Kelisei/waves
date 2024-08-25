package api

import (
	"fmt"
	"net/http"
	"waves/sessions"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Obtener la sesión actual
	session, _ := sessions.Store.Get(r, "session-name")

	// Verificar si el usuario está autenticado
	auth, ok := session.Values["authenticated"].(bool)
	if !ok || !auth {
		http.Error(w, "No estás autenticado", http.StatusUnauthorized)
		return
	}

	// Mostrar un mensaje de bienvenida si está autenticado
	fmt.Fprintln(w, "Bienvenido a la página principal!")
}
