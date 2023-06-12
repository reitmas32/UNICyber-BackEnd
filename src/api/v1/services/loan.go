package services

import (
	"time"

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

func FindLoanByAccountNumber(idStudent uint) (bool, string, models.Loan) {

	var loan models.Loan
	if err := config.DB.First(&loan, "id_student = ? AND sesion_end = ?", idStudent, time.Time{}).Error; err != nil {
		return false, "No Find Loan", loan
	}

	return true, "Find Room", loan
}

func LoanLeaveComputer(idComputer uint) (bool, string, models.Loan) {

	var loan models.Loan
	if err := config.DB.First(&loan, "id_computer = ? AND sesion_end = ?", idComputer, time.Time{}).Error; err != nil {
		return false, "No Find Room", loan
	}
	loan.SesionEnd = time.Now()

	config.DB.Save(&loan)

	return true, "Find Loan", loan
}
