package config

import (
	"os"
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

	// Obtener valores de las variables de entorno
	ENVIRONMENT = getEnv("ENVIRONMENT")
	SECRET_KEY_TOKEN = getEnv("SECRET_KEY_TOKEN")
	UNIACCOUNTS_API_SERVICE_NAME = getEnv("UNIACCOUNTS_API_SERVICE_NAME")
	UNIACCOUNTS_API_KEY = getEnv("UNIACCOUNTS_API_KEY")
	SMTP_USER = getEnv("SMTP_USER")
	SMTP_PASSWORD = getEnv("SMTP_PASSWORD")

}

func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
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
