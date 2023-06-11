package routes

import (
	"fmt"

	"github.com/UNIHacks/UNIAccounts-BackEnd/src/api/v1/views"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/config"
)

func UniversityPrograms() {
	//Get Room
	config.Router.GET(fmt.Sprintf("/api/%s/university-programs", config.API_VERSION), views.UniversityPrograms_GET)
}