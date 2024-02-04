// funcion de operaciones matematicas

package main

import (
	"fmt"
	"net/http"
	"strconv" //convercion de tipo de datos
)

// num12 es un nuevo tipo basado en float64
type num12 float64

var historial []string //slide para almacenar el historial de operaciones

// funciones de operaciones aritmeticas
func suma(a, b num12) num12 {
	return a + b
}

func resta(a, b num12) num12 {
	return a - b
}

func multiplicacion(a, b num12) num12 {
	return a * b
}

func division(a, b num12) (num12, string) {
	if b == 0 {
		return 0, "No se puede dividir por cero"
	}
	return a / b, ""
}

func operacion(w http.ResponseWriter, r *http.Request, op func(num12, num12) num12, operador string) {
	valorStr := r.URL.Query().Get("valor")
	if valorStr == "" {
		http.Error(w, "Se requiere el parámetro 'valor'", http.StatusBadRequest)
		return
	}

	valor, err := strconv.ParseFloat(valorStr, 64)
	if err != nil {
		http.Error(w, "Formato de número inválido", http.StatusBadRequest)
		return
	}

	resultado := op(num12(10), num12(valor)) //esto es una prueba
	fmt.Fprintf(w, "%s: %.2f\n", operador, resultado)

	historial = append(historial, fmt.Sprintf("%s %.2f = %.2f", operador, num12(10), resultado))
}

// realizar funciones y guardar en historial
func Sumar(w http.ResponseWriter, r *http.Request) {
	operacion(w, r, suma, "Suma")
}

func Restar(w http.ResponseWriter, r *http.Request) {
	operacion(w, r, resta, "Resta")
}

func Multiplicar(w http.ResponseWriter, r *http.Request) {
	operacion(w, r, multiplicacion, "Multiplicación")
}

func Dividir(w http.ResponseWriter, r *http.Request) {
	valorStr := r.URL.Query().Get("valor")
	if valorStr == "" {
		http.Error(w, "Se requiere el parámetro 'valor'", http.StatusBadRequest)
		return
	}

	valor, err := strconv.ParseFloat(valorStr, 64)
	if err != nil {
		http.Error(w, "Formato de número inválido", http.StatusBadRequest)
		return
	}

	resultado, mensajeError := division(num12(10), num12(valor))
	if mensajeError != "" {
		http.Error(w, mensajeError, http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "División: %.2f\n", resultado)

	historial = append(historial, fmt.Sprintf("División %.2f = %.2f", num12(10), resultado))
}

// devuelve el historial de operaciones
func HistorialHandler(w http.ResponseWriter, r *http.Request) {
	for _, op := range historial {
		fmt.Fprintln(w, op)
	}
}
