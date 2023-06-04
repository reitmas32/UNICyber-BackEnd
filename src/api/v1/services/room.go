package services

import (
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/api/v1/schemas"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/config"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/models"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/tools"
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

func FindRoom(IdRoom string) (bool, string, models.Room) {

	var room models.Room
	if err := config.DB.First(&room, "id_room = ?", IdRoom).Error; err != nil {
		return false, "No Find Room", room
	}

	return true, "Find Room", room
}

func DeleteRoom(IdRoom string) (bool, string, models.Room) {

	var room models.Room
	if err := config.DB.Delete(&room, "id_room = ?", IdRoom).Error; err != nil {
		return false, "No Find Room", room
	}

	return true, "Delete Room", room
}

func UpdateRoom(IdRoom string, new_compuer schemas.RoomUpdateSchema) (bool, string, models.Room) {

	var room models.Room
	if err := config.DB.First(&room, "id_room = ?", IdRoom).Error; err != nil {
		return false, "No Find Computer", room
	} else {
		room.Name = tools.CopyField(new_compuer.Name, room.Name, "")
		room.Index = tools.CopyField(new_compuer.Index, room.Index, -1)

		config.DB.Save(&room)
	}

	return true, "Change Name Successful", room
}
