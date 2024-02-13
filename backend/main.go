// Entrada de app back
package main

import (
	"fmt"
	"net/http"
)

/*
func HistorialHandler(w http.ResponseWriter, r *http.Request) {
    for _, op := range historial {
        fmt.Fprintln(w, op)
    }
}*/

func main() {

	// Configurar rutas
	http.HandleFunc("/test", handler)
	http.HandleFunc("/sumar", Sumar)
	http.HandleFunc("/restar", Restar)
	http.HandleFunc("/multiplicar", Multiplicar)
	http.HandleFunc("/dividir", Dividir)
	http.HandleFunc("/historial", HistorialHandler)

	//archivos estaticos
	http.HandleFunc("/static/", StaticHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("calculadora-app/public"))))
	// servidor
	port := 8000
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
