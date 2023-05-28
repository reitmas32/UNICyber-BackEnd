package services

import (
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/config"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/models"
)

func CreateRoom(room models.Room) (bool, string) {

	result := config.DB.Create(&room)
	if result.Error != nil {
		return false, result.Error.Error()
	}

	return true, "Create Room Successful"
}
