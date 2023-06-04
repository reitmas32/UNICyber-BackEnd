package config

import "github.com/gin-gonic/gin"

// Configuración de las rutas y controladores de la API
var Router *gin.Engine

func SetupRouter() {
	Router = gin.Default()
}
