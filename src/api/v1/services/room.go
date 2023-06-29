package services

import (
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/api/v1/schemas"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/config"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/models"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/tools"
)

// CreateRoom creates a new room in the database.
// It takes a Room object as input and inserts it into the database.
// It returns a tuple containing a boolean value indicating the success or failure of the operation,
// a string message describing the result, and the created Room object.
//
// Parameters:
//
//	room (models.Room): The Room object to be created.
//
// Returns:
//
//	success (bool): A boolean value indicating the success (true) or failure (false) of the operation.
//	message (string): A string message providing information about the result of the operation.
//	room (models.Room): The created Room object, if the operation was successful.
//	                    Otherwise, it will contain default values for its fields.
//
// Example:
//
//	room := models.Room{
//	    Name:          "Room A",
//	    IdComputerLab: 1,
//	    Index:         1,
//	}
//	success, message, createdRoom := CreateRoom(room)
//	if success {
//	    fmt.Println(message)
//	    fmt.Println("Created Room:", createdRoom)
//	} else {
//	    fmt.Println("Error:", message)
//	}
func CreateRoom(room models.Room) (bool, string, models.Room) {

	result := config.DB.Create(&room)
	if result.Error != nil {
		return false, result.Error.Error(), room
	}

	return true, "Create Room Successful", room
}

// FindRoom retrieves a room from the database based on the provided room ID.
// It returns a tuple containing a boolean value indicating the success or failure of the operation,
// a string message describing the result, and the retrieved Room object.
//
// Parameters:
//
//	id (uint): The ID of the room to retrieve.
//
// Returns:
//
//	success (bool): A boolean value indicating the success (true) or failure (false) of the retrieval operation.
//	message (string): A string message providing information about the result of the retrieval operation.
//	room (models.Room): The retrieved Room object, if found in the database.
//	                   Otherwise, it will contain default values for its fields.
//
// Example:
//
//	success, message, room := FindRoom(1)
//	if success {
//	    fmt.Println(message)
//	    fmt.Println("Found Room:", room)
//	} else {
//	    fmt.Println("Error:", message)
//	}
func FindRoom(id uint) (bool, string, models.Room) {

	var room models.Room
	if err := config.DB.First(&room, "id = ?", id).Error; err != nil {
		return false, "No Find Room", room
	}

	return true, "Find Room", room
}

// DeleteRoom deletes a room from the database based on the provided room ID.
// It returns a tuple containing a boolean value indicating the success or failure of the operation,
// a string message describing the result, and the deleted Room object.
//
// Parameters:
//    id (uint): The ID of the room to delete.
//
// Returns:
//    success (bool): A boolean value indicating the success (true) or failure (false) of the deletion operation.
//    message (string): A string message providing information about the result of the deletion operation.
//    room (models.Room): The deleted Room object, if the deletion operation was successful.
//                       Otherwise, it will contain default values for its fields.
//
// Example:
//    success, message, deletedRoom := DeleteRoom(1)
//    if success {
//        fmt.Println(message)
//        fmt.Println("Deleted Room:", deletedRoom)
//    } else {
//        fmt.Println("Error:", message)
//    }

func DeleteRoom(id uint) (bool, string, models.Room) {

	var room models.Room
	if err := config.DB.Delete(&room, "id = ?", id).Error; err != nil {
		return false, "No Find Room", room
	}

	return true, "Delete Room", room
}

// UpdateRoom updates a room in the database based on the provided room ID and new room data.
// It returns a tuple containing a boolean value indicating the success or failure of the operation,
// a string message describing the result, and the updated Room object.
//
// Parameters:
//
//	id (uint): The ID of the room to update.
//	new_room (schemas.RoomUpdateSchema): The new room data to update.
//
// Returns:
//
//	success (bool): A boolean value indicating the success (true) or failure (false) of the update operation.
//	message (string): A string message providing information about the result of the update operation.
//	room (models.Room): The updated Room object, if the update operation was successful.
//	                   Otherwise, it will contain default values for its fields.
//
// Example:
//
//	newRoom := schemas.RoomUpdateSchema{
//	    Name:  "Updated Room",
//	    Index: 2,
//	}
//	success, message, updatedRoom := UpdateRoom(1, newRoom)
//	if success {
//	    fmt.Println(message)
//	    fmt.Println("Updated Room:", updatedRoom)
//	} else {
//	    fmt.Println("Error:", message)
//	}
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

// FindRoomsOfComputerLab retrieves all rooms associated with a specific computer lab from the database.
// It takes the ID of the computer lab as input and returns a tuple containing a boolean value indicating
// the success or failure of the operation, a string message describing the result, and a slice of Room objects.
//
// Parameters:
//
//	id_computer_lab (uint): The ID of the computer lab to retrieve rooms for.
//
// Returns:
//
//	success (bool): A boolean value indicating the success (true) or failure (false) of the retrieval operation.
//	message (string): A string message providing information about the result of the retrieval operation.
//	rooms ([]models.Room): A slice of Room objects associated with the specified computer lab.
//	                       If no rooms are found or an error occurs, an empty slice will be returned.
//
// Example:
//
//	success, message, rooms := FindRoomsOfComputerLab(1)
//	if success {
//	    fmt.Println(message)
//	    for _, room := range rooms {
//	        fmt.Println("Room:", room)
//	    }
//	} else {
//	    fmt.Println("Error:", message)
//	}
func FindRoomsOfComputerLab(id_computer_lab uint) (bool, string, []models.Room) {

	var rooms []models.Room
	result := config.DB.Where("id_computer_lab = ?", id_computer_lab).Find(&rooms)
	if result.Error != nil {
		return false, result.Error.Error(), rooms
	}
	return true, "Find Rooms by ComputerLabs", rooms
}
