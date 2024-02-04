// Entrada de app back
package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "hola mundo desde el servidor go")
}

func main() {
	// Configurar rutas
	http.HandleFunc("/", handler)

	// servidor
	port := 8080
	fmt.Printf("Servidor escuchando en http://localhost:%d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error:", err)
	}

}
