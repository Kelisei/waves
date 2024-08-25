package sessions

import (
	"github.com/gorilla/sessions"
)

// Crear un almacenamiento de sesiones global
var Store = sessions.NewCookieStore([]byte("super-secret-key"))
