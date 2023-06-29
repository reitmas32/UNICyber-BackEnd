package services

import (
	"time"

	"github.com/UNIHacks/UNIAccounts-BackEnd/src/config"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/models"
)

// CreateLoan creates a loan by persisting the provided Loan object in the database.
// It returns a tuple containing a boolean value indicating the success or failure of the operation,
// a string message describing the result, and the created Loan object.
//
// Parameters:
//
//	loan (Loan): The Loan object representing the loan to be created.
//
// Returns:
//
//	success (bool): A boolean value indicating the success (true) or failure (false) of the create operation.
//	message (string): A string message providing information about the result of the create operation.
//	loan (Loan): The created Loan object, including any modifications made during the create operation
//	             (such as auto-generated IDs or timestamps).
//
// Example:
//
//	loan := Loan{IdStudent: 1, IdComputer: 2, SessionStart: time.Now(), SessionEnd: time.Now().Add(time.Hour)}
//	success, message, createdLoan := CreateLoan(loan)
//	if success {
//	    fmt.Println(message)
//	    fmt.Println("Created Loan:", createdLoan)
//	} else {
//	    fmt.Println("Error:", message)
//	}
func CreateLoan(loan models.Loan) (bool, string, models.Loan) {

	result := config.DB.Create(&loan)
	if result.Error != nil {
		return false, result.Error.Error(), loan
	}

	return true, "Create Loan Successful", loan
}

// FindLoanByAccountNumber retrieves a loan from the database based on the provided student ID.
// It returns a tuple containing a boolean value indicating the success or failure of the operation,
// a string message describing the result, and the retrieved Loan object.
//
// Parameters:
//
//	idStudent (uint): The student ID to retrieve the loan for.
//
// Returns:
//
//	success (bool): A boolean value indicating the success (true) or failure (false) of the retrieval operation.
//	message (string): A string message providing information about the result of the retrieval operation.
//	loan (models.Loan): The retrieved Loan object, if found in the database.
//	                   Otherwise, it will contain default values for its fields.
//
// Example:
//
//	success, message, loan := FindLoanByAccountNumber(1)
//	if success {
//	    fmt.Println(message)
//	    fmt.Println("Found Loan:", loan)
//	} else {
//	    fmt.Println("Error:", message)
//	}
func FindLoanByAccountNumber(idStudent uint) (bool, string, models.Loan) {

	var loan models.Loan
	if err := config.DB.First(&loan, "id_student = ? AND sesion_end = ?", idStudent, time.Time{}).Error; err != nil {
		return false, "No Find Loan", loan
	}

	return true, "Find Room", loan
}

// LoanLeaveComputer marks the end of a loan for a specific computer in the database.
// It updates the session end time of the loan with the current time.
// It returns a tuple containing a boolean value indicating the success or failure of the operation,
// a string message describing the result, and the updated Loan object.
//
// Parameters:
//
//	idComputer (uint): The ID of the computer for which the loan will be ended.
//
// Returns:
//
//	success (bool): A boolean value indicating the success (true) or failure (false) of the operation.
//	message (string): A string message providing information about the result of the operation.
//	loan (models.Loan): The updated Loan object, if found and successfully updated in the database.
//	                   Otherwise, it will contain default values for its fields.
//
// Example:
//
//	success, message, loan := LoanLeaveComputer(1)
//	if success {
//	    fmt.Println(message)
//	    fmt.Println("Updated Loan:", loan)
//	} else {
//	    fmt.Println("Error:", message)
//	}
func LoanLeaveComputer(idComputer uint) (bool, string, models.Loan) {

	var loan models.Loan
	if err := config.DB.First(&loan, "id_computer = ? AND sesion_end = ?", idComputer, time.Time{}).Error; err != nil {
		return false, "No Find Room", loan
	}
	loan.SesionEnd = time.Now()

	config.DB.Save(&loan)

	return true, "Find Loan", loan
}

// FindLoanByIdComputer retrieves a loan from the database based on the provided computer ID.
// It returns a tuple containing a boolean value indicating the success or failure of the operation,
// a string message describing the result, and the retrieved Loan object.
//
// Parameters:
//
//	idComputer (uint): The computer ID to retrieve the loan for.
//
// Returns:
//
//	success (bool): A boolean value indicating the success (true) or failure (false) of the retrieval operation.
//	message (string): A string message providing information about the result of the retrieval operation.
//	loan (models.Loan): The retrieved Loan object, if found in the database.
//	                   Otherwise, it will contain default values for its fields.
//
// Example:
//
//	success, message, loan := FindLoanByIdComputer(1)
//	if success {
//	    fmt.Println(message)
//	    fmt.Println("Found Loan:", loan)
//	} else {
//	    fmt.Println("Error:", message)
//	}
func FindLoanByIdComputer(idComputer uint) (bool, string, models.Loan) {

	var loan models.Loan
	if err := config.DB.First(&loan, "id_computer = ? AND sesion_end = ?", idComputer, time.Time{}).Error; err != nil {
		return false, "No Find Computer", loan
	}

	return true, "Find Computer", loan
}
