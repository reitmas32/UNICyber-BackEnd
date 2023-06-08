package services

import (
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/config"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/models"
)

func CreateLinkAccount(linkAccount models.LinkAccount) (bool, string, models.LinkAccount) {

	result := config.DB.Create(&linkAccount)
	if result.Error != nil {
		return false, result.Error.Error(), linkAccount
	}

	return true, "Create LinkAccount Successful", linkAccount
}

func GetComputerLabs_User(user_name string) (bool, string, []models.LinkAccount) {
	var linksAccount []models.LinkAccount
	result := config.DB.Where("user_name = ?", user_name).Find(&linksAccount)
	if result.Error != nil {
		return false, result.Error.Error(), linksAccount
	}
	return true, "Create LinkAccount Successful", linksAccount

}
