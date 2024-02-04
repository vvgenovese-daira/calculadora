//Rutas y operaciones de la calculadora

package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "hola recibida en handler")
	fmt.Fprint(w, "hola, desde servidor goland")
}
