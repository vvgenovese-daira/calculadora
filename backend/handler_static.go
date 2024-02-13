package main

import (
	"net/http"
	"strings"
)

// manejo de archivos estáticos
func StaticHandler(w http.ResponseWriter, r *http.Request) {
	filePath := strings.TrimPrefix(r.URL.Path, "/static/")
	http.ServeFile(w, r, "calculadora-app/"+filePath)
}
