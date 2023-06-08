package services

import (
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/config"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/models"
)

func CreateLinkAccountCode(linkAccountCode models.LinkAccountCode) (bool, string, models.LinkAccountCode) {

	result := config.DB.Create(&linkAccountCode)
	if result.Error != nil {
		return false, result.Error.Error(), linkAccountCode
	}

	return true, "Create Code Successful", linkAccountCode
}

func FindLinkAccountCode(code string, user_name string) (bool, string, models.LinkAccountCode) {

	var linkAccountCode models.LinkAccountCode
	if err := config.DB.First(&linkAccountCode, "code = ? AND user_name = ?", code, user_name).Error; err != nil {
		return false, "No Exist the ComputerLab", linkAccountCode
	}

	return true, "Create Code Successful", linkAccountCode
}
