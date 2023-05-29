package services

import (
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/config"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/models"
)

func CreateComputer(computer models.Computer) (bool, string) {

	var room models.Room
	if err := config.DB.First(&room, "id_room = ?", computer.IdRoom).Error; err != nil {
		return false, "No existe la Aula"
	}

	result := config.DB.Create(&computer)
	if result.Error != nil {
		return false, result.Error.Error()
	}

	return true, "Create Computer Successful"
}

func FindComputer(IdComputer string) (bool, string, models.Computer) {

	var computer models.Computer
	if err := config.DB.First(&computer, "id_computer = ?", IdComputer).Error; err != nil {
		return false, "No Find Computer", computer
	}

	return true, "Find Computer", computer
}

func DeleteComputer(IdComputer string) (bool, string, models.Computer) {

	var computer models.Computer
	if err := config.DB.Delete(&computer, "id_computer = ?", IdComputer).Error; err != nil {
		return false, "No Find Computer", computer
	}

	return true, "Delete Computer", computer
}
