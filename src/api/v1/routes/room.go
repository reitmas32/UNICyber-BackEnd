package routes

import (
	"fmt"

	"github.com/UNIHacks/UNIAccounts-BackEnd/src/api/v1/views"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/config"
)

func Rooms() {
	//Create Room
	config.Router.POST(fmt.Sprintf("/api/%s/rooms", config.API_VERSION), views.Room_POST)
}
