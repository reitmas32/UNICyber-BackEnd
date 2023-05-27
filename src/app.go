package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/UNIHacks/UNIAccounts-BackEnd/src/api/v1/routes"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/config"
	"github.com/gorilla/mux"
)

func main() {
	config.LoadEnvs()
	log.Println("Run Server on ENVIRONMENT:", config.ENVIRONMENT)
	log.Println("Run Server on PORT:", config.PORT)
	log.Println("Run Server on DEBUG:", config.DEBUG)
	log.Println("Run Server on API_VERSION:", config.API_VERSION)
	log.Println("UNIACCOUNTS_API_SERVICE_NAME:", config.UNIACCOUNTS_API_SERVICE_NAME)
	log.Println("UNIACCOUNTS_API_KEY:", config.UNIACCOUNTS_API_KEY)

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", routes.IndexHandler)
	router.HandleFunc(fmt.Sprintf("/api/%s/signup", config.API_VERSION), routes.SignUp_POST)
	router.HandleFunc(fmt.Sprintf("/api/%s/signin", config.API_VERSION), routes.SignIn_PUT)

	log.Fatal(http.ListenAndServe(fmt.Sprint(":", config.PORT), router))
}
