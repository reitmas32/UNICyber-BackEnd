package routes

import (
	"fmt"

	"github.com/UNIHacks/UNIAccounts-BackEnd/src/api/v1/views"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/config"
)

func StudentRoutes() {
	//Create Computer
	config.Router.POST(fmt.Sprintf("/api/%s/student", config.API_VERSION), views.Student_POST)

	//Get Computer
	config.Router.GET(fmt.Sprintf("/api/%s/student/:id", config.API_VERSION), views.Student_GET)

	//Update Computer
	config.Router.PUT(fmt.Sprintf("/api/%s/student/:id", config.API_VERSION), views.Student_PUT)

	//Delete Computer
	config.Router.DELETE(fmt.Sprintf("/api/%s/student/:id", config.API_VERSION), views.Student_DELETE)

	//Get Student
	config.Router.GET(fmt.Sprintf("/api/%s/students/:account-number", config.API_VERSION), views.StudentByAccountNumber_GET)
}
