package routes

import (
	"fmt"

	"github.com/UNIHacks/UNIAccounts-BackEnd/src/api/v1/views"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/config"
)

func States() {
	//Get Room
	config.Router.GET(fmt.Sprintf("/api/%s/states", config.API_VERSION), views.States_GET)
}
