package services

import (
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/config"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/models"
)

func CreateState(state models.State) (bool, string, models.State) {

	result := config.DB.Create(&state)
	if result.Error != nil {
		return false, result.Error.Error(), state
	}

	return true, "Create State Successful", state
}

func FindState(id uint) (bool, string, models.State) {

	var state models.State
	if err := config.DB.First(&state, "id = ?", id).Error; err != nil {
		return false, "No Find State", state
	}

	return true, "Find State", state
}
