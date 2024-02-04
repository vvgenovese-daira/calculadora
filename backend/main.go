// Entrada de app back
package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "hola recibida en handler")
	fmt.Fprint(w, "hola, desde servidor goland")
}

func main() {
	// Configurar rutas
	http.HandleFunc("/test", handler)

	// servidor
	port := 8000
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
