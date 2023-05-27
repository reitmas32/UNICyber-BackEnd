package routes

import (
	"io/ioutil"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// Hacer una solicitud GET a http://127.0.0.1:5000
	resp, err := http.Get("http://127.0.0.1:5000")
	if err != nil {
		http.Error(w, "Error al obtener el contenido del archivo", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Leer el contenido de la respuesta
	htmlBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error al leer el contenido del archivo", http.StatusInternalServerError)
		return
	}

	// Establecer el encabezado Content-Type a "text/html"
	w.Header().Set("Content-Type", "text/html")

	// Escribir el contenido HTML en el cuerpo de la respuesta
	w.Write(htmlBytes)
}
