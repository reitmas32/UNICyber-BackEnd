package services

import (
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/api/v1/schemas"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/config"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/models"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/tools"
)

func CreateComputer(computer models.Computer) (bool, string, models.Computer) {

	result := config.DB.Create(&computer)
	if result.Error != nil {
		return false, result.Error.Error(), computer
	}

	return true, "Create Computer Successful", computer
}

func FindComputer(id uint) (bool, string, models.Computer) {

	var computer models.Computer
	if err := config.DB.First(&computer, "id = ?", id).Error; err != nil {
		return false, "No Find Computer", computer
	}

	return true, "Find Computer", computer
}

func DeleteComputer(id uint) (bool, string, models.Computer) {

	var computer models.Computer
	if err := config.DB.Delete(&computer, "id = ?", id).Error; err != nil {
		return false, "No Find Computer", computer
	}

	return true, "Delete Computer", computer
}

func UpdateComputer(id uint, new_compuer schemas.ComputerUpdateSchema) (bool, string, models.Computer) {

	var computer models.Computer
	if err := config.DB.First(&computer, "id = ?", id).Error; err != nil {
		return false, "No Find Computer", computer
	} else {
		computer.Name = tools.CopyField(new_compuer.Name, computer.Name, "")
		computer.Type = tools.CopyField(new_compuer.Type, computer.Type, "")
		computer.State = tools.CopyField(new_compuer.State, computer.State, "")
		computer.Message = tools.CopyField(new_compuer.Message, computer.Message, "")
		computer.Pos_x = tools.CopyField(new_compuer.Pos_x, computer.Pos_x, 0.0)
		computer.Pos_y = tools.CopyField(new_compuer.Pos_y, computer.Pos_y, 0.0)

		config.DB.Save(&computer)
	}

	return true, "Change Name Successful", computer
}

func FindComputerOfRoom(id_room uint) (bool, string, []models.Computer) {

	var computers []models.Computer
	result := config.DB.Where("id_room = ?", id_room).Find(&computers)
	if result.Error != nil {
		return false, result.Error.Error(), computers
	}
	return true, "Find Computers by Room", computers
}

func SetStateComputer(id uint, id_state uint) (bool, string, models.Computer) {

	var computer models.Computer
	if err := config.DB.First(&computer, "id = ?", id).Error; err != nil {
		return false, "No Find Computer", computer
	} else {

		computer.IdState = tools.CopyField(id_state, computer.IdState, uint(0))

		config.DB.Save(&computer)
	}

	return true, "Set State Successful", computer
}
