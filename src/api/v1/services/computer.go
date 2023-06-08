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

func FindComputer(id string) (bool, string, models.Computer) {

	var computer models.Computer
	if err := config.DB.First(&computer, "id = ?", id).Error; err != nil {
		return false, "No Find Computer", computer
	}

	return true, "Find Computer", computer
}

func DeleteComputer(id string) (bool, string, models.Computer) {

	var computer models.Computer
	if err := config.DB.Delete(&computer, "id = ?", id).Error; err != nil {
		return false, "No Find Computer", computer
	}

	return true, "Delete Computer", computer
}

func UpdateComputer(id string, new_compuer schemas.ComputerUpdateSchema) (bool, string, models.Computer) {

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

/*
func copyFieldFloat32(src float32, des float32, default_value float32) float32 {
	if src != default_value {
		des = src
	}
	return des
}
*/
