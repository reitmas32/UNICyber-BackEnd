package services

import (
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/config"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/models"
)

// CreateLinkAccountCode creates a link account code by persisting the provided LinkAccountCode object in the database.
// It returns a tuple containing a boolean value indicating the success or failure of the operation,
// a string message describing the result, and the created LinkAccountCode object.
//
// Parameters:
//
//	linkAccountCode (LinkAccountCode): The LinkAccountCode object representing the link account code to be created.
//
// Returns:
//
//	success (bool): A boolean value indicating the success (true) or failure (false) of the create operation.
//	message (string): A string message providing information about the result of the create operation.
//	linkAccountCode (LinkAccountCode): The created LinkAccountCode object, including any modifications made during the create operation
//	                                   (such as auto-generated IDs or timestamps).
//
// Example:
//
//	linkCode := LinkAccountCode{UserName: "john_doe", Code: "569778", IdComputerLab: 1}
//	success, message, createdLinkCode := CreateLinkAccountCode(linkCode)
//	if success {
//	    fmt.Println(message)
//	    fmt.Println("Created LinkAccountCode:", createdLinkCode)
//	} else {
//	    fmt.Println("Error:", message)
//	}
func CreateLinkAccountCode(linkAccountCode models.LinkAccountCode) (bool, string, models.LinkAccountCode) {

	result := config.DB.Create(&linkAccountCode)
	if result.Error != nil {
		return false, result.Error.Error(), linkAccountCode
	}

	return true, "Create Code Successful", linkAccountCode
}

// FindLinkAccountCode retrieves a link account code from the database based on the provided code and user_name.
// It returns a tuple containing a boolean value indicating the success or failure of the operation,
// a string message describing the result, and the retrieved LinkAccountCode object.
//
// Parameters:
//
//	code (string): The code of the link account code to retrieve from the database.
//	user_name (string): The user name associated with the link account code.
//
// Returns:
//
//	success (bool): A boolean value indicating the success (true) or failure (false) of the retrieval operation.
//	message (string): A string message providing information about the result of the retrieval operation.
//	linkAccountCode (models.LinkAccountCode): The retrieved LinkAccountCode object, if found in the database.
//	                                          Otherwise, it will contain default values for its fields.
//
// Example:
//
//	success, message, linkCode := FindLinkAccountCode("569778", "john_doe")
//	if success {
//	    fmt.Println(message)
//	    fmt.Println("Found LinkAccountCode:", linkCode)
//	} else {
//	    fmt.Println("Error:", message)
//	}
func FindLinkAccountCode(code string, user_name string) (bool, string, models.LinkAccountCode) {

	var linkAccountCode models.LinkAccountCode
	if err := config.DB.First(&linkAccountCode, "code = ? AND user_name = ?", code, user_name).Error; err != nil {
		return false, "No Exist the ComputerLab", linkAccountCode
	}

	return true, "Create Code Successful", linkAccountCode
}
