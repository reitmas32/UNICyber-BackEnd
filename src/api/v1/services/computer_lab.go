package services

import (
	"fmt"

	"github.com/UNIHacks/UNIAccounts-BackEnd/src/api/v1/schemas"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/config"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/models"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/tools"
)

// CreateComputerLab creates a computer lab by persisting the provided ComputerLab object in the database.
// It returns a tuple containing a boolean value indicating the success or failure of the operation,
// a string message describing the result, and the updated ComputerLab object.
//
// Parameters:
//
//	computerLab (ComputerLab): The ComputerLab object representing the computer lab to be created.
//
// Returns:
//
//	success (bool): A boolean value indicating the success (true) or failure (false) of the create operation.
//	message (string): A string message providing information about the result of the create operation.
//	updatedComputerLab (ComputerLab): The updated ComputerLab object, which includes any modifications made during the create operation.
//
// Example:
//
//	computerLab := ComputerLab{Name: "Lab 1", Description: "Lab for programming"}
//	success, message, updatedComputerLab := CreateComputerLab(computerLab)
//	if success {
//	    fmt.Println(message)
//	    fmt.Println("Created ComputerLab:", updatedComputerLab)
//	} else {
//	    fmt.Println("Error:", message)
//	}
func CreateComputerLab(computerLab models.ComputerLab) (bool, string, models.ComputerLab) {

	result := config.DB.Create(&computerLab)
	if result.Error != nil {
		return false, result.Error.Error(), computerLab
	}

	return true, "Create ComputerLab Successful", computerLab
}

// FindComputerLab retrieves a computer lab from the database based on the provided ID.
// It returns a tuple containing a boolean value indicating the success or failure of the operation,
// a string message describing the result, and the retrieved ComputerLab object.
//
// Parameters:
//
//	id (uint): The ID of the computer lab to retrieve from the database.
//
// Returns:
//
//	success (bool): A boolean value indicating the success (true) or failure (false) of the find operation.
//	message (string): A string message providing information about the result of the find operation.
//	computerLab (models.ComputerLab): The retrieved ComputerLab object, if found in the database.
//	                                 Otherwise, it will contain default values for its fields.
//
// Example:
//
//	success, message, computerLab := FindComputerLab(1)
//	if success {
//	    fmt.Println(message)
//	    fmt.Println("Found ComputerLab:", computerLab)
//	} else {
//	    fmt.Println("Error:", message)
//	}
func FindComputerLab(id uint) (bool, string, models.ComputerLab) {

	var computerLab models.ComputerLab
	if err := config.DB.First(&computerLab, "id = ?", id).Error; err != nil {
		return false, fmt.Sprint("No Find Computer Lab: ", err), computerLab
	}

	return true, "Find Computer Lab", computerLab
}

// DeleteComputerLab deletes a computer lab from the database based on the provided ID.
// It returns a tuple containing a boolean value indicating the success or failure of the operation,
// a string message describing the result, and the deleted ComputerLab object.
//
// Parameters:
//
//	id (uint): The ID of the computer lab to delete from the database.
//
// Returns:
//
//	success (bool): A boolean value indicating the success (true) or failure (false) of the delete operation.
//	message (string): A string message providing information about the result of the delete operation.
//	computerLab (models.ComputerLab): The deleted ComputerLab object, if found and successfully deleted from the database.
//	                                 Otherwise, it will contain default values for its fields.
//
// Example:
//
//	success, message, computerLab := DeleteComputerLab(1)
//	if success {
//	    fmt.Println(message)
//	    fmt.Println("Deleted ComputerLab:", computerLab)
//	} else {
//	    fmt.Println("Error:", message)
//	}
func DeleteComputerLab(id uint) (bool, string, models.ComputerLab) {

	var computerLab models.ComputerLab
	if err := config.DB.Delete(&computerLab, "id = ?", id).Error; err != nil {
		return false, "No Find Computer Lab", computerLab
	}

	return true, "Delete Computer Lab", computerLab
}

// UpdateComputerLab updates the information of a computer lab in the database based on the provided ID and new_computer parameters.
// It returns a tuple containing a boolean value indicating the success or failure of the operation,
// a string message describing the result, and the updated ComputerLab object.
//
// Parameters:
//
//	id (uint): The ID of the computer lab to update in the database.
//	new_computer (schemas.ComputerLabUpdateSchema): The updated information for the computer lab.
//
// Returns:
//
//	success (bool): A boolean value indicating the success (true) or failure (false) of the update operation.
//	message (string): A string message providing information about the result of the update operation.
//	computerLab (models.ComputerLab): The updated ComputerLab object, if found and successfully updated in the database.
//	                                 Otherwise, it will contain default values for its fields.
//
// Example:
//
//	newComputer := schemas.ComputerLabUpdateSchema{Name: "Lab 1 Updated", Description: "Updated description"}
//	success, message, updatedComputerLab := UpdateComputerLab(1, newComputer)
//	if success {
//	    fmt.Println(message)
//	    fmt.Println("Updated ComputerLab:", updatedComputerLab)
//	} else {
//	    fmt.Println("Error:", message)
//	}
func UpdateComputerLab(id uint, new_compuer schemas.ComputerLabUpdateSchema) (bool, string, models.ComputerLab) {

	var computerLab models.ComputerLab
	if err := config.DB.First(&computerLab, "id = ?", id).Error; err != nil {
		return false, "No Find Computer Lab", computerLab
	} else {
		computerLab.Name = tools.CopyField(new_compuer.Name, computerLab.Name, "")
		computerLab.Description = tools.CopyField(new_compuer.Description, computerLab.Description, "")

		config.DB.Save(&computerLab)
	}

	return true, "Update ComputerLab Successful", computerLab
}

// GetComputerLabs retrieves all computer labs from the database.
// It returns a tuple containing a boolean value indicating the success or failure of the operation,
// a string message describing the result, and a slice of ComputerLab objects representing the retrieved computer labs.
//
// Returns:
//
//	success (bool): A boolean value indicating the success (true) or failure (false) of the retrieval operation.
//	message (string): A string message providing information about the result of the retrieval operation.
//	computerLabs ([]models.ComputerLab): A slice of ComputerLab objects representing the retrieved computer labs.
//	                                     If no computer labs are found, the slice will be empty.
//
// Example:
//
//	success, message, computerLabs := GetComputerLabs()
//	if success {
//	    fmt.Println(message)
//	    for _, lab := range computerLabs {
//	        fmt.Println("ComputerLab:", lab)
//	    }
//	} else {
//	    fmt.Println("Error:", message)
//	}
func GetComputerLabs() (bool, string, []models.ComputerLab) {
	var computerLabs []models.ComputerLab
	if err := config.DB.Find(&computerLabs).Error; err != nil {
		return false, fmt.Sprint("No se encontraron Computer Labs: ", err), computerLabs
	}

	return true, "Computer Labs encontrados", computerLabs
}

// GetComputerLabs_Limit retrieves a specified number of computer labs from the database based on the provided limit.
// It returns a tuple containing a boolean value indicating the success or failure of the operation,
// a string message describing the result, and a slice of ComputerLab objects representing the retrieved computer labs.
//
// Parameters:
//
//	limit (int): The maximum number of computer labs to retrieve from the database.
//
// Returns:
//
//	success (bool): A boolean value indicating the success (true) or failure (false) of the retrieval operation.
//	message (string): A string message providing information about the result of the retrieval operation.
//	computerLabs ([]models.ComputerLab): A slice of ComputerLab objects representing the retrieved computer labs.
//	                                     If no computer labs are found, the slice will be empty.
//
// Example:
//
//	success, message, computerLabs := GetComputerLabs_Limit(5)
//	if success {
//	    fmt.Println(message)
//	    for _, lab := range computerLabs {
//	        fmt.Println("ComputerLab:", lab)
//	    }
//	} else {
//	    fmt.Println("Error:", message)
//	}
func GetComputerLabs_Limit(limit int) (bool, string, []models.ComputerLab) {
	var computerLabs []models.ComputerLab
	if err := config.DB.Limit(limit).Find(&computerLabs).Error; err != nil {
		return false, fmt.Sprint("No se encontraron Computer Labs: ", err), computerLabs
	}

	return true, "Computer Labs encontrados", computerLabs
}
