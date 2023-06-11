package services

import (
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/api/v1/schemas"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/config"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/models"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/tools"
)

func CreateStudent(student models.Student) (bool, string, models.Student) {

	result := config.DB.Create(&student)
	if result.Error != nil {
		return false, result.Error.Error(), student
	}

	return true, "Create Student Successful", student
}

func FindStudent(id uint) (bool, string, models.Student) {

	var student models.Student
	if err := config.DB.First(&student, "id = ?", id).Error; err != nil {
		return false, "No Find Student", student
	}

	return true, "Find Student", student
}

func DeleteStudent(id uint) (bool, string, models.Student) {

	var student models.Student
	if err := config.DB.Delete(&student, "id = ?", id).Error; err != nil {
		return false, "No Find Student", student
	}

	return true, "Delete Student", student
}

func UpdateStudent(id uint, new_student schemas.StudentUpdateSchema) (bool, string, models.Student) {

	var student models.Student
	if err := config.DB.First(&student, "id = ?", id).Error; err != nil {
		return false, "No Find Student", student
	} else {
		student.Name = tools.CopyField(new_student.Name, student.Name, "")
		student.LastName = tools.CopyField(new_student.LastName, student.LastName, "")
		student.Email = tools.CopyField(new_student.Email, student.Email, "")
		student.UniversityProgram = tools.CopyField(new_student.UniversityProgram, student.UniversityProgram, "")
		student.AccountNumber = tools.CopyField(new_student.AccountNumber, student.AccountNumber, "")
		student.Semester = tools.CopyField(new_student.Semester, student.Semester, 0)
		student.IdUniversityProgram = tools.CopyField(new_student.IdUniversityProgram, student.IdUniversityProgram, 0)

		config.DB.Save(&student)
	}

	return true, "Update Student Successful", student
}

func FindStudentByAccountNumber(accountNumber string) (bool, string, models.Student) {

	var student models.Student
	if err := config.DB.First(&student, "account_number = ?", accountNumber).Error; err != nil {
		return false, "No Find Student", student
	}

	return true, "Find Student", student
}
