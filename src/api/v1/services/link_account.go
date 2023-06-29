package services

import (
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/config"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/models"
)

// CreateLinkAccount creates a link account by persisting the provided LinkAccount object in the database.
// It returns a tuple containing a boolean value indicating the success or failure of the operation,
// a string message describing the result, and the created LinkAccount object.
//
// Parameters:
//
//	linkAccount (LinkAccount): The LinkAccount object representing the link account to be created.
//
// Returns:
//
//	success (bool): A boolean value indicating the success (true) or failure (false) of the create operation.
//	message (string): A string message providing information about the result of the create operation.
//	linkAccount (LinkAccount): The created LinkAccount object, including any modifications made during the create operation
//	                           (such as auto-generated IDs or timestamps).
//
// Example:
//
//	link := LinkAccount{UserName: "john_doe", IdComputerLab: 1}
//	success, message, createdLink := CreateLinkAccount(link)
//	if success {
//	    fmt.Println(message)
//	    fmt.Println("Created LinkAccount:", createdLink)
//	} else {
//	    fmt.Println("Error:", message)
//	}
func CreateLinkAccount(linkAccount models.LinkAccount) (bool, string, models.LinkAccount) {

	result := config.DB.Create(&linkAccount)
	if result.Error != nil {
		return false, result.Error.Error(), linkAccount
	}

	return true, "Create LinkAccount Successful", linkAccount
}

// GetComputerLabs_User retrieves all computer labs associated with a specific user from the database.
// It returns a tuple containing a boolean value indicating the success or failure of the operation,
// a string message describing the result, and a slice of LinkAccount objects representing the retrieved links between the user and computer labs.
//
// Parameters:
//
//	user_name (string): The user name to retrieve computer labs for.
//
// Returns:
//
//	success (bool): A boolean value indicating the success (true) or failure (false) of the retrieval operation.
//	message (string): A string message providing information about the result of the retrieval operation.
//	linksAccount ([]models.LinkAccount): A slice of LinkAccount objects representing the retrieved links between the user and computer labs.
//	                                     If no links are found, the slice will be empty.
//
// Example:
//
//	success, message, links := GetComputerLabs_User("john_doe")
//	if success {
//	    fmt.Println(message)
//	    for _, link := range links {
//	        fmt.Println("Link Account:", link)
//	    }
//	} else {
//	    fmt.Println("Error:", message)
//	}
func GetComputerLabs_User(user_name string) (bool, string, []models.LinkAccount) {
	var linksAccount []models.LinkAccount
	result := config.DB.Where("user_name = ?", user_name).Find(&linksAccount)
	if result.Error != nil {
		return false, result.Error.Error(), linksAccount
	}
	return true, "Find ComputerLabs by User", linksAccount

}
