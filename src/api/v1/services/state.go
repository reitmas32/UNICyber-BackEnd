package services

import (
	"fmt"

	"github.com/UNIHacks/UNIAccounts-BackEnd/src/config"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/models"
)

// CreateState creates a new state in the database.
// It takes a State object as input and returns a tuple containing a boolean value indicating
// the success or failure of the operation, a string message describing the result, and the created State object.
//
// Parameters:
//
//	state (models.State): The State object to be created in the database.
//
// Returns:
//
//	success (bool): A boolean value indicating the success (true) or failure (false) of the creation operation.
//	message (string): A string message providing information about the result of the creation operation.
//	state (models.State): The State object that has been created. If an error occurs during creation, this will be the original state object passed as input.
//
// Example:
//
//	newState := models.State{
//	    Name:  "New State",
//	    Image: "state.png",
//	}
//	success, message, createdState := CreateState(newState)
//	if success {
//	    fmt.Println("State created successfully:", createdState)
//	} else {
//	    fmt.Println("Error:", message)
//	}
func CreateState(state models.State) (bool, string, models.State) {

	result := config.DB.Create(&state)
	if result.Error != nil {
		return false, result.Error.Error(), state
	}

	return true, "Create State Successful", state
}

// FindState retrieves a state from the database based on its ID.
// It takes an ID as input and returns a tuple containing a boolean value indicating
// the success or failure of the operation, a string message describing the result, and the retrieved State object.
//
// Parameters:
//
//	id (uint): The ID of the state to be retrieved from the database.
//
// Returns:
//
//	success (bool): A boolean value indicating the success (true) or failure (false) of the retrieval operation.
//	message (string): A string message providing information about the result of the retrieval operation.
//	state (models.State): The State object that has been retrieved. If the state is not found, this will be an empty state object.
//
// Example:
//
//	success, message, foundState := FindState(1)
//	if success {
//	    fmt.Println("State found:", foundState)
//	} else {
//	    fmt.Println("Error:", message)
//	}
func FindState(id uint) (bool, string, models.State) {

	var state models.State
	if err := config.DB.First(&state, "id = ?", id).Error; err != nil {
		return false, "No Find State", state
	}

	return true, "Find State", state
}

// GetStates retrieves all states from the database.
// It returns a tuple containing a boolean value indicating the success or failure of the operation,
// a string message describing the result, and a slice of State objects representing the retrieved states.
//
// Returns:
//
//	success (bool): A boolean value indicating the success (true) or failure (false) of the retrieval operation.
//	message (string): A string message providing information about the result of the retrieval operation.
//	states ([]models.State): A slice of State objects representing the retrieved states. If no states are found, this slice will be empty.
//
// Example:
//
//	success, message, foundStates := GetStates()
//	if success {
//	    for _, state := range foundStates {
//	        fmt.Println("State:", state)
//	    }
//	} else {
//	    fmt.Println("Error:", message)
//	}
func GetStates() (bool, string, []models.State) {

	var states []models.State
	if err := config.DB.Find(&states).Error; err != nil {
		return false, fmt.Sprint("No se encontraron States: ", err), states
	}

	return true, "States encontrados", states
}
