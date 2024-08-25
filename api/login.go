package api

import (
	"fmt"
	"net/http"
	"waves/sessions"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Obtener la sesión actual
	session, _ := sessions.Store.Get(r, "session-name")

	// Autenticación exitosa (esto es solo un ejemplo, deberías verificar el usuario/contraseña)
	session.Values["authenticated"] = true
	session.Save(r, w)

	fmt.Fprintln(w, "Has iniciado sesión exitosamente")
}
