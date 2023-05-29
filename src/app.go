package main

import (
	"fmt"
	"log"

	"github.com/UNIHacks/UNIAccounts-BackEnd/src/api/v1/routes"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/config"

	_ "github.com/mattn/go-sqlite3"
)

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

	config.Router.GET("/", routes.IndexHandler)
	routes.SignIn()
	routes.SignUp()
	routes.Rooms()
	routes.Computer()

	// Escucha en el puerto 8080
	config.Router.Run(fmt.Sprintf(":%d", config.PORT))
}
