package services

import (
	"fmt"

	"github.com/UNIHacks/UNIAccounts-BackEnd/src/api/v1/schemas"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/config"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/models"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/tools"
)

func CreateComputerLab(computerLab models.ComputerLab) (bool, string, models.ComputerLab) {

	result := config.DB.Create(&computerLab)
	if result.Error != nil {
		return false, result.Error.Error(), computerLab
	}

	return true, "Create ComputerLab Successful", computerLab
}

func FindComputerLab(id uint) (bool, string, models.ComputerLab) {

	var computerLab models.ComputerLab
	if err := config.DB.First(&computerLab, "id = ?", id).Error; err != nil {
		return false, fmt.Sprint("No Find Computer Lab: ", err), computerLab
	}

	return true, "Find Computer Lab", computerLab
}

func DeleteComputerLab(id uint) (bool, string, models.ComputerLab) {

	var computerLab models.ComputerLab
	if err := config.DB.Delete(&computerLab, "id = ?", id).Error; err != nil {
		return false, "No Find Computer Lab", computerLab
	}

	return true, "Delete Computer Lab", computerLab
}

func UpdateComputerLab(id uint, new_compuer schemas.ComputerLabUpdateSchema) (bool, string, models.ComputerLab) {

	var computerLab models.ComputerLab
	if err := config.DB.First(&computerLab, "id = ?", id).Error; err != nil {
		return false, "No Find Computer Lab", computerLab
	} else {
		computerLab.Name = tools.CopyField(new_compuer.Name, computerLab.Name, "")
		computerLab.Description = tools.CopyField(new_compuer.Description, computerLab.Description, "")

		config.DB.Save(&computerLab)
	}

	return true, "Update ComputerLab Successful", computerLab
}

func GetComputerLabs() (bool, string, []models.ComputerLab) {
	var computerLabs []models.ComputerLab
	if err := config.DB.Find(&computerLabs).Error; err != nil {
		return false, fmt.Sprint("No se encontraron Computer Labs: ", err), computerLabs
	}

	return true, "Computer Labs encontrados", computerLabs
}

func GetComputerLabs_Limit(limit int) (bool, string, []models.ComputerLab) {
	var computerLabs []models.ComputerLab
	if err := config.DB.Limit(limit).Find(&computerLabs).Error; err != nil {
		return false, fmt.Sprint("No se encontraron Computer Labs: ", err), computerLabs
	}

	return true, "Computer Labs encontrados", computerLabs
}
