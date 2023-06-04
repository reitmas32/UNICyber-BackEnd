package services

import (
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/config"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/models"
)

func CreateLinkAccount(linkAccount models.LinkAccount) (bool, string) {

	result := config.DB.Create(&linkAccount)
	if result.Error != nil {
		return false, result.Error.Error()
	}

	return true, "Create LinkAccount Successful"
}
