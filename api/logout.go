package api

import (
	"fmt"
	"net/http"
	"waves/sessions"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// Obtener la sesión actual
	session, _ := sessions.Store.Get(r, "session-name")

	// Eliminar la autenticación
	session.Values["authenticated"] = false
	session.Save(r, w)

	fmt.Fprintln(w, "Has cerrado sesión exitosamente")
}
