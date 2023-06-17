package main

import (
	"fmt"
	"log"

	_ "github.com/UNIHacks/UNIAccounts-BackEnd/src/docs"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/models"

	//swaggerFiles "github.com/swaggo/files"
	//ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/UNIHacks/UNIAccounts-BackEnd/src/api/v1/routes"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/api/v1/services"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/config"

	_ "github.com/mattn/go-sqlite3"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
)

func CreateStates() {
	createState(1, "Disponible", "https://em-content.zobj.net/source/microsoft-teams/337/desktop-computer_1f5a5-fe0f.png")
	createState(2, "Descompuesto", "https://raw.githubusercontent.com/reitmas32/unica_cybercoffee/main/public/assets/reparacion.png")
	createState(3, "Mantenimiento", "https://raw.githubusercontent.com/reitmas32/unica_cybercoffee/main/public/assets/mantenimiento.png")
	createState(4, "Proyecto", "https://raw.githubusercontent.com/reitmas32/unica_cybercoffee/main/public/assets/proyecto.png")
	createState(5, "Reparación", "https://raw.githubusercontent.com/reitmas32/unica_cybercoffee/main/public/assets/reparacion.png")
	createState(6, "Loan", "https://raw.githubusercontent.com/reitmas32/unica_cybercoffee/main/public/assets/proyecto.png")

}

func createState(id uint, name string, img string) {
	result, _, _ := services.FindState(id)
	if !result {
		state := models.State{
			Name:  name,
			Image: img,
		}
		services.CreateState(state)
	}
}

func CreateUniversityPrograms() {
	createUniversityProgram(1, "Ingeniería Aeroespacial")
	createUniversityProgram(2, "Ingeniería Civil")
	createUniversityProgram(3, "Ingeniería Geomática")
	createUniversityProgram(4, "Ingeniería Ambiental")
	createUniversityProgram(5, "Ingeniería Geofísica")
	createUniversityProgram(6, "Ingeniería Geológica")
	createUniversityProgram(7, "Ingeniería Petrolera")
	createUniversityProgram(8, "Ingeniería de Minas y Metalurgia")
	createUniversityProgram(9, "Ingeniería en Computación")
	createUniversityProgram(10, "Ingeniería Eléctrica Electrónica")
	createUniversityProgram(11, "Ingeniería en Telecomunicaciones")
	createUniversityProgram(12, "Ingeniería Mecánica")
	createUniversityProgram(13, "Ingeniería Industrial")
	createUniversityProgram(14, "Ingeniería Mecatrónica")
	createUniversityProgram(15, "Ingeniería en Sistemas Biomédicos")
	createUniversityProgram(16, "Posgrado")
	createUniversityProgram(17, "Intercambio")

}

func createUniversityProgram(id uint, name string) {
	result, _, _ := services.FindUniversityProgram(id)
	if !result {
		universityProgram := models.UniversityProgram{
			Name: name,
		}
		services.CreateUniversityProgram(universityProgram)
	}
}

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

	CreateStates()
	CreateUniversityPrograms()

	config.SetupRouter()

	// docs route on Mode Debug
	//config.Router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	config.Router.GET("/", routes.IndexHandler)
	routes.SignIn()
	routes.SignUp()
	routes.LinkAccountWithComputerLab()
	routes.Rooms()
	routes.Computer()
	routes.ComputerLab()
	routes.StudentRoutes()
	routes.LoanRoutes()
	routes.States()
	routes.UniversityPrograms()

	// Escucha en el puerto 8080
	config.Router.Run(fmt.Sprintf(":%d", config.PORT))
}
