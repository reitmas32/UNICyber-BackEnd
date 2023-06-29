package config

import "github.com/gin-gonic/gin"

var Router *gin.Engine

func SetupRouter() {
	Router = gin.Default()
}
