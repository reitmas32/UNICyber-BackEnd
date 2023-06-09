package routes

import (
	"fmt"

	"github.com/UNIHacks/UNIAccounts-BackEnd/src/api/v1/views"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/config"
)

func Rooms() {
	//Create Room
	config.Router.POST(fmt.Sprintf("/api/%s/room", config.API_VERSION), views.Room_POST)

	//Get Room
	config.Router.GET(fmt.Sprintf("/api/%s/room/:id", config.API_VERSION), views.Room_GET)

	//Update Room
	config.Router.PUT(fmt.Sprintf("/api/%s/room/:id", config.API_VERSION), views.Room_PUT)

	//Delete Room
	config.Router.DELETE(fmt.Sprintf("/api/%s/room/:id", config.API_VERSION), views.Room_DELETE)

	//Get Rooms of ComputerLab
	config.Router.GET(fmt.Sprintf("/api/%s/rooms/:id-computer-lab", config.API_VERSION), views.RoomsOfCompuer_GET)
}
