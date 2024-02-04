// Entrada de app back
package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Configurar rutas
	http.HandleFunc("/test", handler)
	http.HandleFunc("/sumar", Sumar)
	http.HandleFunc("/restar", Restar)
	http.HandleFunc("/multiplicar", Multiplicar)
	http.HandleFunc("/dividir", Dividir)
	http.HandleFunc("/historial", HistorialHandler)

	// servidor
	port := 8000
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
