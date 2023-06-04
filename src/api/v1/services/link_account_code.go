package services

import (
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/config"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/models"
)

func CreateLinkAccountCode(linkAccountCode models.LinkAccountCode) (bool, string) {

	var computerLab models.ComputerLab
	if err := config.DB.First(&computerLab, "id_computer_lab = ?", linkAccountCode.IdComputerLab).Error; err != nil {
		return false, "No Exist the ComputerLab"
	}

	result := config.DB.Create(&linkAccountCode)
	if result.Error != nil {
		return false, result.Error.Error()
	}

	return true, "Create Code Successful"
}

func FindLinkAccountCode(code string, user_name string) (bool, string, models.LinkAccountCode) {

	var linkAccountCode models.LinkAccountCode
	if err := config.DB.First(&linkAccountCode, "code = ? AND user_name = ?", code, user_name).Error; err != nil {
		return false, "No Exist the ComputerLab", linkAccountCode
	}

	return true, "Create Code Successful", linkAccountCode
}
