package services

import (
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/api/v1/schemas"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/config"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/models"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/tools"
)

func CreateRoom(room models.Room) (bool, string, models.Room) {

	result := config.DB.Create(&room)
	if result.Error != nil {
		return false, result.Error.Error(), room
	}

	return true, "Create Room Successful", room
}

func FindRoom(id uint) (bool, string, models.Room) {

	var room models.Room
	if err := config.DB.First(&room, "id = ?", id).Error; err != nil {
		return false, "No Find Room", room
	}

	return true, "Find Room", room
}

func DeleteRoom(id uint) (bool, string, models.Room) {

	var room models.Room
	if err := config.DB.Delete(&room, "id = ?", id).Error; err != nil {
		return false, "No Find Room", room
	}

	return true, "Delete Room", room
}

func UpdateRoom(id uint, new_room schemas.RoomUpdateSchema) (bool, string, models.Room) {

	var room models.Room
	if err := config.DB.First(&room, "id = ?", id).Error; err != nil {
		return false, "No Find Room", room
	} else {
		room.Name = tools.CopyField(new_room.Name, room.Name, "")
		room.Index = tools.CopyField(new_room.Index, room.Index, -1)

		config.DB.Save(&room)
	}

	return true, "Update Room Successful", room
}

func FindRoomsOfComputerLab(id_computer_lab uint) (bool, string, []models.Room) {

	var rooms []models.Room
	result := config.DB.Where("id_computer_lab = ?", id_computer_lab).Find(&rooms)
	if result.Error != nil {
		return false, result.Error.Error(), rooms
	}
	return true, "Find Rooms by ComputerLabs", rooms
}
