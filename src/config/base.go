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
	SMTP_HOST                    = "smtp.office365.com"
	SMTP_PORT                    = 587
	SMTP_USER                    string
	SMTP_PASSWORD                string
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
	SMTP_USER = getEnv("SMTP_USER")
	SMTP_PASSWORD = getEnv("SMTP_PASSWORD")

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
		case "SMTP_USER":
			return "user@mail.com"
		case "SMTP_PASSWORD":
			return "password"
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
