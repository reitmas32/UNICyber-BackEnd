package services

import (
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/config"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/models"
)

func CreateRoom(room models.Room) (bool, string) {

	var computerLab models.ComputerLab
	if err := config.DB.First(&computerLab, "id_computer_lab = ?", room.IdComputerLab).Error; err != nil {
		return false, "No Exist the ComputerLab"
	}

	result := config.DB.Create(&room)
	if result.Error != nil {
		return false, result.Error.Error()
	}

	return true, "Create Room Successful"
}
