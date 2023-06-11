package views

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/UNIHacks/UNIAccounts-BackEnd/src/api/v1/schemas"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/api/v1/services"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// @Summary add a new item of the students
// @ID create-student
// @Tags Students
// @Produce json
// @Param data body schemas.StudentCreateSchema true "Schema by Create New Student"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /api/v1/student [post]
func Student_POST(c *gin.Context) {

	responseCreateStudent := models.Response{
		Message: "No Create Student",
		Success: false,
		Data:    "{}",
	}

	// Decodificar el objeto JSON recibido
	var studentCreateSchema schemas.StudentCreateSchema
	if err := c.ShouldBindWith(&studentCreateSchema, binding.JSON); err != nil {
		responseCreateStudent := models.Response{
			Message: "Error to Get Content JSON",
			Success: false,
			Data:    strings.Split(err.Error(), "\n"),
		}
		c.Header("Content-Type", "application/json")
		c.JSON(200, responseCreateStudent)
		return
	}

	student := models.Student{
		Name:              studentCreateSchema.Name,
		LastName:          studentCreateSchema.LastName,
		UniversityProgram: studentCreateSchema.UniversityProgram,
		Email:             studentCreateSchema.Email,
		AccountNumber:     studentCreateSchema.AccountNumber,
		Semester:          studentCreateSchema.Semester,
	}

	result, message, new_student := services.CreateStudent(student)

	if result {

		responseCreateStudent = models.Response{
			Message: message,
			Success: result,
			Data:    new_student,
		}
	} else {
		responseCreateStudent = models.Response{
			Message: message,
			Success: responseCreateStudent.Success,
			Data:    "{}",
		}
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, responseCreateStudent)
}

// @Summary get a item of the students
// @ID get-student
// @Tags Students
// @Produce json
// @Param id path string true "ID of Student"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /api/v1/student [get]
func Student_GET(c *gin.Context) {

	id, _ := (strconv.ParseUint(c.Param("id"), 10, 32))

	result, message, student := services.FindStudent(uint(id))

	responseGetStudent := models.Response{
		Message: message,
		Success: result,
		Data:    student,
	}

	if !result {
		responseGetStudent.Data = "{}"
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, responseGetStudent)
}

// @Summary delete a item of the students
// @ID delete-student
// @Tags Students
// @Produce json
// @Param id path string true "ID of Student"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /api/v1/student [delete]
func Student_DELETE(c *gin.Context) {

	id, _ := (strconv.ParseUint(c.Param("id"), 10, 32))

	result, message, student := services.DeleteStudent(uint(id))

	responseDeleteStudent := models.Response{
		Message: message,
		Success: result,
		Data:    student,
	}

	if !result {
		responseDeleteStudent.Data = "{}"
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, responseDeleteStudent)
}

// @Summary update a item of the students
// @ID put-student
// @Tags Students
// @Produce json
// @Param id path string true "ID of Students"
// @Param data body schemas.StudentUpdateSchema true "Schema by Update New Student"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /api/v1/student [put]
func Student_PUT(c *gin.Context) {

	// Decodificar el objeto JSON recibido
	var studentUpdateSchema schemas.StudentUpdateSchema
	err := json.NewDecoder(c.Request.Body).Decode(&studentUpdateSchema)
	if err != nil {
		responseUpdateStudent := models.Response{
			Message: "Error to Get Content JSON",
			Success: false,
			Data:    "{}",
		}
		c.Header("Content-Type", "application/json")
		c.JSON(200, responseUpdateStudent)
		return
	}

	id, _ := (strconv.ParseUint(c.Param("id"), 10, 32))

	result, message, student := services.FindStudent(uint(id))

	if result {
		result, message, student = services.UpdateStudent(uint(id), studentUpdateSchema)
	}

	responseUpdateStudent := models.Response{
		Message: message,
		Success: result,
		Data:    student,
	}

	if !result {
		responseUpdateStudent.Data = "{}"
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, responseUpdateStudent)
}

// @Summary get a item of the students
// @ID get-student
// @Tags Students
// @Produce json
// @Param id path string true "ID of Student"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /api/v1/student [get]
func StudentByAccountNumber_GET(c *gin.Context) {

	accountNumber := c.Param("account-number")

	result, message, student := services.FindStudentByAccountNumber(accountNumber)

	responseGetStudent := models.Response{
		Message: message,
		Success: result,
		Data:    student,
	}

	if !result {
		responseGetStudent.Data = "{}"
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, responseGetStudent)
}
