package services

import (
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/api/v1/schemas"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/config"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/models"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/tools"
)

func CreateComputerLab(computerLab models.ComputerLab) (bool, string) {

	result := config.DB.Create(&computerLab)
	if result.Error != nil {
		return false, result.Error.Error()
	}

	return true, "Create ComputerLab Successful"
}

func FindComputerLab(IdComputerLab string) (bool, string, models.ComputerLab) {

	var computerLab models.ComputerLab
	if err := config.DB.First(&computerLab, "id_computer_lab = ?", IdComputerLab).Error; err != nil {
		return false, "No Find Computer", computerLab
	}

	return true, "Find Computer Lab", computerLab
}

func DeleteComputerLab(IdComputerLab string) (bool, string, models.ComputerLab) {

	var computerLab models.ComputerLab
	if err := config.DB.Delete(&computerLab, "id_computer_lab = ?", IdComputerLab).Error; err != nil {
		return false, "No Find Computer Lab", computerLab
	}

	return true, "Delete Computer Lab", computerLab
}

func UpdateComputerLab(IdComputerLab string, new_compuer schemas.ComputerLabUpdateSchema) (bool, string, models.ComputerLab) {

	var computerLab models.ComputerLab
	if err := config.DB.First(&computerLab, "id_computer_lab = ?", IdComputerLab).Error; err != nil {
		return false, "No Find Computer Lab", computerLab
	} else {
		computerLab.Name = tools.CopyField(new_compuer.Name, computerLab.Name, "")
		computerLab.Description = tools.CopyField(new_compuer.Description, computerLab.Description, "")

		config.DB.Save(&computerLab)
	}

	return true, "Update ComputerLab Successful", computerLab
}
