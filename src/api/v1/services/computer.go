package services

import (
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/api/v1/schemas"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/config"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/models"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/tools"
)

////////////////////
//// CRUD
////////////////////

// CreateComputer creates a computer by persisting the provided Computer object in the database.
// It returns a tuple containing a boolean value indicating the success or failure of the operation,
// a string message describing the result, and the created Computer object.
//
// Parameters:
//
//	computer (Computer): The Computer object representing the computer to be created.
//
// Returns:
//
//	success (bool): A boolean value indicating the success (true) or failure (false) of the create operation.
//	message (string): A string message providing information about the result of the create operation.
//	computer (Computer): The created Computer object, including any modifications made during the create operation
//	                    (such as auto-generated IDs or timestamps).
//
// Example:
//
//	computer := Computer{Name: "Computer 1", State: "Active", Type: "Desktop"}
//	success, message, createdComputer := CreateComputer(computer)
//	if success {
//	    fmt.Println(message)
//	    fmt.Println("Created Computer:", createdComputer)
//	} else {
//	    fmt.Println("Error:", message)
//	}
func CreateComputer(computer models.Computer) (bool, string, models.Computer) {

	result := config.DB.Create(&computer)
	if result.Error != nil {
		return false, result.Error.Error(), computer
	}

	return true, "Create Computer Successful", computer
}

// FindComputer retrieves a computer from the database based on the provided ID.
// It returns a tuple containing a boolean value indicating the success or failure of the operation,
// a string message describing the result, and the retrieved Computer object.
//
// Parameters:
//
//	id (uint): The ID of the computer to retrieve from the database.
//
// Returns:
//
//	success (bool): A boolean value indicating the success (true) or failure (false) of the find operation.
//	message (string): A string message providing information about the result of the find operation.
//	computer (models.Computer): The retrieved Computer object, if found in the database.
//	                           Otherwise, it will contain default values for its fields.
//
// Example:
//
//	success, message, computer := FindComputer(1)
//	if success {
//	    fmt.Println(message)
//	    fmt.Println("Found Computer:", computer)
//	} else {
//	    fmt.Println("Error:", message)
//	}
func FindComputer(id uint) (bool, string, models.Computer) {

	var computer models.Computer
	if err := config.DB.First(&computer, "id = ?", id).Error; err != nil {
		return false, "No Find Computer", computer
	}

	return true, "Find Computer", computer
}

// DeleteComputer deletes a computer from the database based on the provided ID.
// It returns a tuple containing a boolean value indicating the success or failure of the operation,
// a string message describing the result, and the deleted Computer object.
//
// Parameters:
//
//	id (uint): The ID of the computer to delete from the database.
//
// Returns:
//
//	success (bool): A boolean value indicating the success (true) or failure (false) of the delete operation.
//	message (string): A string message providing information about the result of the delete operation.
//	computer (models.Computer): The deleted Computer object, if found and successfully deleted from the database.
//	                           Otherwise, it will contain default values for its fields.
//
// Example:
//
//	success, message, computer := DeleteComputer(1)
//	if success {
//	    fmt.Println(message)
//	    fmt.Println("Deleted Computer:", computer)
//	} else {
//	    fmt.Println("Error:", message)
//	}
func DeleteComputer(id uint) (bool, string, models.Computer) {

	var computer models.Computer
	if err := config.DB.Delete(&computer, "id = ?", id).Error; err != nil {
		return false, "No Find Computer", computer
	}

	return true, "Delete Computer", computer
}

// UpdateComputer updates the information of a computer in the database based on the provided ID and new_computer parameters.
// It returns a tuple containing a boolean value indicating the success or failure of the operation,
// a string message describing the result, and the updated Computer object.
//
// Parameters:
//
//	id (uint): The ID of the computer to update in the database.
//	new_computer (schemas.ComputerUpdateSchema): The updated information for the computer.
//
// Returns:
//
//	success (bool): A boolean value indicating the success (true) or failure (false) of the update operation.
//	message (string): A string message providing information about the result of the update operation.
//	computer (models.Computer): The updated Computer object, if found and successfully updated in the database.
//	                           Otherwise, it will contain default values for its fields.
//
// Example:
//
//	newComputer := schemas.ComputerUpdateSchema{Name: "Computer 1 Updated", Type: "Laptop", State: "Inactive", Message: "Updated message", Pos_x: 1.5, Pos_y: 2.0}
//	success, message, updatedComputer := UpdateComputer(1, newComputer)
//	if success {
//	    fmt.Println(message)
//	    fmt.Println("Updated Computer:", updatedComputer)
//	} else {
//	    fmt.Println("Error:", message)
//	}
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

	return true, "Update Computer Successful", computer
}

///////////////////
//// Operations
///////////////////

// FindComputerOfRoom retrieves all computers associated with a specific room from the database.
// It returns a tuple containing a boolean value indicating the success or failure of the operation,
// a string message describing the result, and a slice of Computer objects representing the retrieved computers.
//
// Parameters:
//
//	id_room (uint): The ID of the room to retrieve computers from.
//
// Returns:
//
//	success (bool): A boolean value indicating the success (true) or failure (false) of the retrieval operation.
//	message (string): A string message providing information about the result of the retrieval operation.
//	computers ([]models.Computer): A slice of Computer objects representing the retrieved computers.
//	                              If no computers are found, the slice will be empty.
//
// Example:
//
//	success, message, computers := FindComputerOfRoom(1)
//	if success {
//	    fmt.Println(message)
//	    for _, comp := range computers {
//	        fmt.Println("Computer:", comp)
//	    }
//	} else {
//	    fmt.Println("Error:", message)
//	}
func FindComputerOfRoom(id_room uint) (bool, string, []models.Computer) {

	var computers []models.Computer
	result := config.DB.Where("id_room = ?", id_room).Find(&computers)
	if result.Error != nil {
		return false, result.Error.Error(), computers
	}
	return true, "Find Computers by Room", computers
}

// SetStateComputer sets the state of a computer in the database based on the provided ID and id_state parameters.
// It returns a tuple containing a boolean value indicating the success or failure of the operation,
// a string message describing the result, and the updated Computer object.
//
// Parameters:
//
//	id (uint): The ID of the computer to update in the database.
//	id_state (uint): The ID of the state to set for the computer.
//
// Returns:
//
//	success (bool): A boolean value indicating the success (true) or failure (false) of the state update operation.
//	message (string): A string message providing information about the result of the state update operation.
//	computer (models.Computer): The updated Computer object, if found and successfully updated in the database.
//	                           Otherwise, it will contain default values for its fields.
//
// Example:
//
//	success, message, updatedComputer := SetStateComputer(1, 2)
//	if success {
//	    fmt.Println(message)
//	    fmt.Println("Updated Computer:", updatedComputer)
//	} else {
//	    fmt.Println("Error:", message)
//	}
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
