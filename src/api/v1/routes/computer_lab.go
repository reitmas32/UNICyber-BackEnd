package routes

import (
	"fmt"

	"github.com/UNIHacks/UNIAccounts-BackEnd/src/api/v1/views"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/config"
)

func ComputerLab() {
	//Create Computer
	config.Router.POST(fmt.Sprintf("/api/%s/computer-lab", config.API_VERSION), views.Computer_POST)

	//Get Computer
	config.Router.GET(fmt.Sprintf("/api/%s/computer-lab/:id-computer", config.API_VERSION), views.Computer_GET)

	//Update Computer
	config.Router.PUT(fmt.Sprintf("/api/%s/computer-lab/:id-computer", config.API_VERSION), views.Computer_PUT)

	//Delete Computer
	config.Router.DELETE(fmt.Sprintf("/api/%s/computer-lab/:id-computer", config.API_VERSION), views.Computer_DELETE)
}