package services

import (
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/config"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/models"
)

func CreateStudent(student models.Student) (bool, string) {

	result := config.DB.Create(&student)
	if result.Error != nil {
		return false, result.Error.Error()
	}

	return true, "Create Student Successful"
}
