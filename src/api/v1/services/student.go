package services

import (
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/api/v1/schemas"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/config"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/models"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/tools"
)

func CreateStudent(student models.Student) (bool, string) {

	result := config.DB.Create(&student)
	if result.Error != nil {
		return false, result.Error.Error()
	}

	return true, "Create Student Successful"
}

func FindStudent(IdStudent string) (bool, string, models.Student) {

	var student models.Student
	if err := config.DB.First(&student, "id_student = ?", IdStudent).Error; err != nil {
		return false, "No Find Student", student
	}

	return true, "Find Student", student
}

func DeleteStudent(IdStudent string) (bool, string, models.Student) {

	var student models.Student
	if err := config.DB.Delete(&student, "id_student = ?", IdStudent).Error; err != nil {
		return false, "No Find Student", student
	}

	return true, "Find Student", student
}

func UpdateStudent(IdStudent string, new_student schemas.StudentUpdateSchema) (bool, string, models.Student) {

	var student models.Student
	if err := config.DB.First(&student, "id_student = ?", IdStudent).Error; err != nil {
		return false, "No Find Computer", student
	} else {
		student.Name = tools.CopyField(new_student.Name, student.Name, "")
		student.LastName = tools.CopyField(new_student.LastName, student.LastName, "")
		student.Email = tools.CopyField(new_student.Email, student.Email, "")
		student.UniversityProgram = tools.CopyField(new_student.UniversityProgram, student.UniversityProgram, "")
		student.AccountNumber = tools.CopyField(new_student.AccountNumber, student.AccountNumber, "")
		student.Semester = tools.CopyField(new_student.Semester, student.Semester, 0)

		config.DB.Save(&student)
	}

	return true, "Change Name Successful", student
}
