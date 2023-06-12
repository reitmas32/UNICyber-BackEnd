package views

import (
	"strings"
	"time"

	"github.com/UNIHacks/UNIAccounts-BackEnd/src/api/v1/schemas"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/api/v1/services"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// @Summary add a new item of the computers
// @ID create-computer
// @Tags Computers
// @Produce json
// @Param data body schemas.loanCreateSchema true "Schema by Create New Computer"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /api/v1/computer [post]
func Loan_POST(c *gin.Context) {

	responseCreateLoan := models.Response{
		Message: "No Create Computer",
		Success: false,
		Data:    "{}",
	}

	// Decodificar el objeto JSON recibido
	var loanCreateSchema schemas.LoanCreateSchema
	if err := c.ShouldBindWith(&loanCreateSchema, binding.JSON); err != nil {
		responseCreateLoan := models.Response{
			Message: "Error to Get Content JSON",
			Success: false,
			Data:    strings.Split(err.Error(), "\n"),
		}
		c.Header("Content-Type", "application/json")
		c.JSON(200, responseCreateLoan)
		return
	}

	exitsStudent, _, _ := services.FindStudent(loanCreateSchema.IdStudent)
	exitsComputer, _, _ := services.FindComputer(loanCreateSchema.IdComputer)

	if !exitsComputer {
		responseCreateLoan.Message = "No Exist the Computer"
		c.Header("Content-Type", "application/json")
		c.JSON(200, responseCreateLoan)
		return
	}

	if !exitsStudent {
		responseCreateLoan.Message = "No Exist the Student"
		c.Header("Content-Type", "application/json")
		c.JSON(200, responseCreateLoan)
		return
	}

	loan := models.Loan{
		IdComputer:  loanCreateSchema.IdComputer,
		IdStudent:   loanCreateSchema.IdStudent,
		SesionStart: time.Now(),
	}

	result, message, new_loan := services.CreateLoan(loan)

	if result {

		responseCreateLoan = models.Response{
			Message: message,
			Success: result,
			Data:    new_loan,
		}
	} else {
		responseCreateLoan = models.Response{
			Message: message,
			Success: responseCreateLoan.Success,
			Data:    nil,
		}
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, responseCreateLoan)
}

// @Summary add a new item of the computers
// @ID create-computer
// @Tags Computers
// @Produce json
// @Param data body schemas.loanCreateSchema true "Schema by Create New Computer"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /api/v1/computer [post]
func LoanByAccountNumber_POST(c *gin.Context) {

	responseCreateLoan := models.Response{
		Message: "No Create Computer",
		Success: false,
		Data:    "{}",
	}

	// Decodificar el objeto JSON recibido
	var loanCreateSchema schemas.LoanCreateByAccountNumberSchema
	if err := c.ShouldBindWith(&loanCreateSchema, binding.JSON); err != nil {
		responseCreateLoan := models.Response{
			Message: "Error to Get Content JSON",
			Success: false,
			Data:    strings.Split(err.Error(), "\n"),
		}
		c.Header("Content-Type", "application/json")
		c.JSON(200, responseCreateLoan)
		return
	}

	exitsStudent, _, student := services.FindStudentByAccountNumber(loanCreateSchema.AccountNumber)
	exitsComputer, _, _ := services.FindComputer(loanCreateSchema.IdComputer)

	if !exitsComputer {
		responseCreateLoan.Message = "No Exist the Computer"
		c.Header("Content-Type", "application/json")
		c.JSON(200, responseCreateLoan)
		return
	}

	if !exitsStudent {
		responseCreateLoan.Message = "No Exist the Student"
		c.Header("Content-Type", "application/json")
		c.JSON(200, responseCreateLoan)
		return
	}

	loan := models.Loan{
		IdComputer:  loanCreateSchema.IdComputer,
		IdStudent:   student.ID,
		SesionStart: time.Now(),
	}

	existLoan, _, old_loan := services.FindLoanByAccountNumber(student.ID)

	if existLoan {
		responseCreateLoan.Message = "Session Start on other computer"
		responseCreateLoan.Data = old_loan
		c.Header("Content-Type", "application/json")
		c.JSON(200, responseCreateLoan)
		return
	}
	result, message, new_loan := services.CreateLoan(loan)

	if result {

		responseCreateLoan = models.Response{
			Message: message,
			Success: result,
			Data:    new_loan,
		}
	} else {
		responseCreateLoan = models.Response{
			Message: message,
			Success: responseCreateLoan.Success,
			Data:    nil,
		}
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, responseCreateLoan)
}
