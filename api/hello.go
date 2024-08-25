package api

import (
	"fmt"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	// Mostrar un mensaje de bienvenida si está autenticado
	fmt.Fprintln(w, "Bienvenido a la página principal!")
}
