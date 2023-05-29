package routes

import (
	"fmt"

	"github.com/UNIHacks/UNIAccounts-BackEnd/src/api/v1/views"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/config"
)

func Computer() {
	//Create Room
	config.Router.POST(fmt.Sprintf("/api/%s/computer", config.API_VERSION), views.Computer_POST)

	//Get Room
	config.Router.GET(fmt.Sprintf("/api/%s/computer/:id-computer", config.API_VERSION), views.Computer_GET)
}
