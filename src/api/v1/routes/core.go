package routes

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	// Hacer una solicitud GET a http://unicappaccountsdev.pythonanywhere.com
	resp, err := http.Get("http://unicappaccountsdev.pythonanywhere.com")
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		c.Writer.Write([]byte("Error al obtener el contenido del archivo"))
		return
	}
	defer resp.Body.Close()

	// Leer el contenido de la respuesta
	htmlBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		c.Writer.Write([]byte("Error al obtener el contenido del archivo"))
		return
	}

	c.Writer.Write(htmlBytes)
}
