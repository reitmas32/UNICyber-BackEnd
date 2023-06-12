package services

import (
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/config"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/models"
)

func CreateLoan(loan models.Loan) (bool, string, models.Loan) {

	result := config.DB.Create(&loan)
	if result.Error != nil {
		return false, result.Error.Error(), loan
	}

	return true, "Create Loan Successful", loan
}
