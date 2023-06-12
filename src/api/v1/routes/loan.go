package routes

import (
	"fmt"

	"github.com/UNIHacks/UNIAccounts-BackEnd/src/api/v1/views"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/config"
)

func LoanRoutes() {
	//Create Computer
	config.Router.POST(fmt.Sprintf("/api/%s/loan-computer", config.API_VERSION), views.Loan_POST)

	//Get Computer
	//config.Router.GET(fmt.Sprintf("/api/%s/loan-computer/:id", config.API_VERSION), views.Student_GET)

	//Update Computer
	//config.Router.PUT(fmt.Sprintf("/api/%s/loan-computer/:id", config.API_VERSION), views.Student_PUT)
}
