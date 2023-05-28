package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateStudent(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusAccepted)
	c.Writer.Write([]byte("Error al obtener el contenido del archivo"))
}

func UpdateStudent(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusAccepted)
	c.Writer.Write([]byte("Error al obtener el contenido del archivo"))
}
func DeleteStudent(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusAccepted)
	c.Writer.Write([]byte("Error al obtener el contenido del archivo"))
}

func GetStudent(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusAccepted)
	c.Writer.Write([]byte("Error al obtener el contenido del archivo"))
}
