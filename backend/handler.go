//Rutas y operaciones de la calculadora

package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hola, desde servidor goland")
}
