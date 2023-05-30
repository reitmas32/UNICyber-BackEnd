package views

import (
	"strings"

	"github.com/UNIHacks/UNIAccounts-BackEnd/src/api/v1/schemas"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/api/v1/services"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
)

func Student_POST(c *gin.Context) {

	responseCreateStudent := map[string]interface{}{
		"Message": "No Create Student",
		"Data":    "{}",
		"Success": false,
	}

	// Decodificar el objeto JSON recibido
	var studentCreateSchema schemas.StudentCreateSchema
	if err := c.ShouldBindWith(&studentCreateSchema, binding.JSON); err != nil {
		responseCreateStudent := map[string]interface{}{
			"Message": "Error to Get Content JSON",
			"Data":    strings.Split(err.Error(), "\n"),
			"Success": false,
		}
		c.Header("Content-Type", "application/json")
		c.JSON(200, responseCreateStudent)
		return
	}

	student := models.Student{
		IdStudent:         uuid.New().String(),
		Name:              studentCreateSchema.Name,
		LastName:          studentCreateSchema.LastName,
		UniversityProgram: studentCreateSchema.UniversityProgram,
		Email:             studentCreateSchema.Email,
		AccountNumber:     studentCreateSchema.AccountNumber,
		Semester:          studentCreateSchema.Semester,
	}

	result, message := services.CreateStudent(student)

	if result {

		responseCreateStudent = map[string]interface{}{
			"Message": message,
			"Success": result,
			"Data":    student,
		}
	} else {
		responseCreateStudent = map[string]interface{}{
			"Message": message,
			"Success": responseCreateStudent["Success"],
			"Data":    student,
		}
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, responseCreateStudent)
}
