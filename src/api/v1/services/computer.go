package services

import (
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/api/v1/schemas"
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

func UpdateComputer(IdComputer string, new_compuer schemas.ComputerUpdateSchema) (bool, string, models.Computer) {

	var computer models.Computer
	if err := config.DB.First(&computer, "id_computer = ?", IdComputer).Error; err != nil {
		return false, "No Find Computer", computer
	} else {
		computer.Name = copyFieldString(new_compuer.Name, computer.Name, "")
		computer.Type = copyFieldString(new_compuer.Type, computer.Type, "")
		computer.State = copyFieldString(new_compuer.State, computer.State, "")
		computer.Message = copyFieldString(new_compuer.Message, computer.Message, "")

		config.DB.Save(&computer)
	}

	return true, "Change Name Successful", computer
}

func copyFieldString(src string, des string, default_value string) string {
	if src != default_value {
		des = src
	}
	return des
}

/*
func copyFieldFloat32(src float32, des float32, default_value float32) float32 {
	if src != default_value {
		des = src
	}
	return des
}
*/
