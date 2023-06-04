package main

import (
	"fmt"
	"log"

	_ "github.com/UNIHacks/UNIAccounts-BackEnd/src/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/UNIHacks/UNIAccounts-BackEnd/src/api/v1/routes"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/config"

	_ "github.com/mattn/go-sqlite3"
)

// @title UNICyber-API
// @version 1.0
// @description This is a API by System UNICyber|SISEC https://github.com/reitmas32/UNICyber-BackEnd

// @contact.name Oswaldo Rafael Zamora Ramirez
// @contact.url https://github.com/reitmas32
// @contact.email rafa.zamora.rals@gmail.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:3000
// @BasePath /
// @query.collection.format multi
func main() {

	config.LoadEnvs()

	log.Println("Run Server on ENVIRONMENT:", config.ENVIRONMENT)
	log.Println("Run Server on PORT:", config.PORT)
	log.Println("Run Server on DEBUG:", config.DEBUG)
	log.Println("Run Server on API_VERSION:", config.API_VERSION)
	log.Println("UNIACCOUNTS_API_SERVICE_NAME:", config.UNIACCOUNTS_API_SERVICE_NAME)
	log.Println("UNIACCOUNTS_API_KEY:", config.UNIACCOUNTS_API_KEY)

	// Configuración de la conexión a la base de datos MySQL
	config.SetupDB()
	// Genera las tablas en la base de datos
	config.MigrateDB()

	config.SetupRouter()

	// docs route on Mode Debug
	config.Router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	config.Router.GET("/", routes.IndexHandler)
	routes.SignIn()
	routes.SignUp()
	routes.LinkAccountWithComputerLab()
	routes.Rooms()
	routes.Computer()
	routes.ComputerLab()
	routes.StudentRoutes()

	// Escucha en el puerto 8080
	config.Router.Run(fmt.Sprintf(":%d", config.PORT))
}
