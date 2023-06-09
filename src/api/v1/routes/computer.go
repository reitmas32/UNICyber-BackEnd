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
	config.Router.GET(fmt.Sprintf("/api/%s/computer/:id", config.API_VERSION), views.Computer_GET)

	//Update Computer
	config.Router.PUT(fmt.Sprintf("/api/%s/computer/:id", config.API_VERSION), views.Computer_PUT)

	//Delete Computer
	config.Router.DELETE(fmt.Sprintf("/api/%s/computer/:id", config.API_VERSION), views.Computer_DELETE)

	//Get Computers of Room
	config.Router.GET(fmt.Sprintf("/api/%s/computers/:id-room", config.API_VERSION), views.ComputersOfRoom_GET)

	//Set State of Computer
	config.Router.PUT(fmt.Sprintf("/api/%s/computer-set-state/:id", config.API_VERSION), views.ComputerSetState_PUT)
}
