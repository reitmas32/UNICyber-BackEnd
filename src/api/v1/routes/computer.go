package routes

import (
	"fmt"

	"github.com/UNIHacks/UNIAccounts-BackEnd/src/api/v1/views"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/config"
)

func Computer() {
	//Create Computer
	config.Router.POST(fmt.Sprintf("/api/%s/computer", config.API_VERSION), views.Computer_POST)

	//Get Computer
	config.Router.GET(fmt.Sprintf("/api/%s/computer/:id-computer", config.API_VERSION), views.Computer_GET)

	//Delete Computer
	config.Router.DELETE(fmt.Sprintf("/api/%s/computer/:id-computer", config.API_VERSION), views.Computer_DELETE)
}
