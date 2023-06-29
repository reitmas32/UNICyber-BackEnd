package services

import (
	"fmt"

	"github.com/UNIHacks/UNIAccounts-BackEnd/src/config"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/models"
)

// CreateUniversityProgram creates a new university program in the database.
// It takes a UniversityProgram object as input and returns a tuple containing a boolean value indicating the success or failure of the operation,
// a string message describing the result, and the created UniversityProgram object (if available).
//
// Parameters:
//
//	universityProgram (models.UniversityProgram): The UniversityProgram object to be created.
//
// Returns:
//
//	success (bool): A boolean value indicating the success (true) or failure (false) of the create operation.
//	message (string): A string message providing information about the result of the create operation.
//	universityProgram (models.UniversityProgram): The created UniversityProgram object, or an empty UniversityProgram object if not created.
//
// Example:
//
//	newUniversityProgram := models.UniversityProgram{Name: "Computer Science"}
//	success, message, createdProgram := CreateUniversityProgram(newUniversityProgram)
//	if success {
//	    fmt.Println("University program created:", createdProgram)
//	} else {
//	    fmt.Println("Error:", message)
//	}
func CreateUniversityProgram(universityProgram models.UniversityProgram) (bool, string, models.UniversityProgram) {

	result := config.DB.Create(&universityProgram)
	if result.Error != nil {
		return false, result.Error.Error(), universityProgram
	}

	return true, "Create UniversityProgram Successful", universityProgram
}

// FindUniversityProgram finds a university program by its ID in the database.
// It takes an ID parameter and returns a tuple containing a boolean value indicating the success or failure of the operation,
// a string message describing the result, and the found UniversityProgram object (if available).
//
// Parameters:
//
//	id (uint): The ID of the university program to find.
//
// Returns:
//
//	success (bool): A boolean value indicating the success (true) or failure (false) of the find operation.
//	message (string): A string message providing information about the result of the find operation.
//	universityProgram (models.UniversityProgram): The found UniversityProgram object, or an empty UniversityProgram object if not found.
//
// Example:
//
//	programID := uint(1)
//	success, message, foundProgram := FindUniversityProgram(programID)
//	if success {
//	    fmt.Println("University program found:", foundProgram)
//	} else {
//	    fmt.Println("Error:", message)
//	}
func FindUniversityProgram(id uint) (bool, string, models.UniversityProgram) {

	var universityProgram models.UniversityProgram
	if err := config.DB.First(&universityProgram, "id = ?", id).Error; err != nil {
		return false, "No Find UniversityProgram", universityProgram
	}

	return true, "Find UniversityProgram", universityProgram
}

// GetUniversityPrograms retrieves all university programs from the database.
// It returns a tuple containing a boolean value indicating the success or failure of the operation,
// a string message describing the result, and a slice of UniversityProgram objects representing the retrieved programs.
//
// Returns:
//
//	success (bool): A boolean value indicating the success (true) or failure (false) of the retrieval operation.
//	message (string): A string message providing information about the result of the retrieval operation.
//	universityPrograms ([]models.UniversityProgram): A slice of UniversityProgram objects representing the retrieved programs, or an empty slice if none were found.
//
// Example:
//
//	success, message, programs := GetUniversityPrograms()
//	if success {
//	    fmt.Println("University programs found:", programs)
//	} else {
//	    fmt.Println("Error:", message)
//	}
func GetUniversityPrograms() (bool, string, []models.UniversityProgram) {

	var universityPrograms []models.UniversityProgram
	if err := config.DB.Find(&universityPrograms).Error; err != nil {
		return false, fmt.Sprint("No se encontraron UniversityPrograms: ", err), universityPrograms
	}

	return true, "UniversityPrograms encontrados", universityPrograms
}
