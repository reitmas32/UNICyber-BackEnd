package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

var (
	ENVIRONMENT      string
	SECRET_KEY_TOKEN string
	API_KEYS         interface{}
	API_VERSION      = "v1"
	UNIACCOUNT_NAME  = "uniaccount"
	TOKEN_TYPE       = "JWT"

	PORT  = 3000
	DEBUG = true
	HOST  = "0.0.0.0"

	//Credentials and Config UNIAccounts
	UNIACCOUNTS_API_VERSION      = "v1"
	UNIACCOUNTS_API_PORT         = 5000
	UNIACCOUNTS_API_KEY          string
	UNIACCOUNTS_API_SERVICE_NAME string
)

func LoadEnvs() {
	// Cargar variables de entorno desde el archivo .env
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatal("Error al cargar el archivo .env:", err)
	}

	// Obtener valores de las variables de entorno
	ENVIRONMENT = getEnv("ENVIRONMENT")
	SECRET_KEY_TOKEN = getEnv("SECRET_KEY_TOKEN")
	UNIACCOUNTS_API_SERVICE_NAME = getEnv("UNIACCOUNTS_API_SERVICE_NAME")
	UNIACCOUNTS_API_KEY = getEnv("UNIACCOUNTS_API_KEY")

}

// getEnv obtiene el valor de una variable de entorno o devuelve un valor predeterminado si no está definida
func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		// Valor predeterminado en caso de que la variable de entorno no esté definida
		// Puedes ajustar esto según tus necesidades
		switch key {
		case "ENVIRONMENT":
			return "development"
		case "SECRET_KEY_TOKEN":
			return "default_secret_key"
		case "UNIACCOUNTS_API_KEY":
			return "api_key"
		case "UNIACCOUNTS_API_SERVICE_NAME":
			return "service"
		default:
			return ""
		}
	}
	return value
}

// getBaseDir obtiene el directorio base del programa
func getBaseDir() string {
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Dir(ex)
}
