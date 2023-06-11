package services

import (
	"fmt"

	"github.com/UNIHacks/UNIAccounts-BackEnd/src/config"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/models"
)

func CreateUniversityProgram(universityProgram models.UniversityProgram) (bool, string, models.UniversityProgram) {

	result := config.DB.Create(&universityProgram)
	if result.Error != nil {
		return false, result.Error.Error(), universityProgram
	}

	return true, "Create UniversityProgram Successful", universityProgram
}

func FindUniversityProgram(id uint) (bool, string, models.UniversityProgram) {

	var universityProgram models.UniversityProgram
	if err := config.DB.First(&universityProgram, "id = ?", id).Error; err != nil {
		return false, "No Find UniversityProgram", universityProgram
	}

	return true, "Find UniversityProgram", universityProgram
}

func GetUniversityPrograms() (bool, string, []models.UniversityProgram) {

	var universityPrograms []models.UniversityProgram
	if err := config.DB.Find(&universityPrograms).Error; err != nil {
		return false, fmt.Sprint("No se encontraron UniversityPrograms: ", err), universityPrograms
	}

	return true, "UniversityPrograms encontrados", universityPrograms
}
