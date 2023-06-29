package services

import (
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/api/v1/schemas"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/config"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/models"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/tools"
)

// CreateStudent creates a new student in the database.
// It takes a Student object as input and returns a tuple containing a boolean value indicating the success or failure of the operation,
// a string message describing the result, and the created Student object.
//
// Parameters:
//
//	student (models.Student): The Student object containing the information of the student to be created.
//
// Returns:
//
//	success (bool): A boolean value indicating the success (true) or failure (false) of the creation operation.
//	message (string): A string message providing information about the result of the creation operation.
//	student (models.Student): The created Student object with the assigned ID and other generated fields.
//
// Example:
//
//	newStudent := models.Student{
//	    Name:                "John",
//	    LastName:            "Doe",
//	    Email:               "johndoe@example.com",
//	    IdUniversityProgram: 123,
//	    AccountNumber:       "ABCD1234",
//	    Semester:            3,
//	}
//	success, message, createdStudent := CreateStudent(newStudent)
//	if success {
//	    fmt.Println("Student created:", createdStudent)
//	} else {
//	    fmt.Println("Error:", message)
//	}
func CreateStudent(student models.Student) (bool, string, models.Student) {

	result := config.DB.Create(&student)
	if result.Error != nil {
		return false, result.Error.Error(), student
	}

	return true, "Create Student Successful", student
}

// FindStudent finds a student in the database by their ID.
// It takes the ID of the student as input and returns a tuple containing a boolean value indicating the success or failure of the operation,
// a string message describing the result, and the found Student object (if available).
//
// Parameters:
//
//	id (uint): The ID of the student to be found.
//
// Returns:
//
//	success (bool): A boolean value indicating the success (true) or failure (false) of the find operation.
//	message (string): A string message providing information about the result of the find operation.
//	student (models.Student): The found Student object, or an empty Student object if not found.
//
// Example:
//
//	success, message, foundStudent := FindStudent(1)
//	if success {
//	    fmt.Println("Student found:", foundStudent)
//	} else {
//	    fmt.Println("Error:", message)
//	}
func FindStudent(id uint) (bool, string, models.Student) {

	var student models.Student
	if err := config.DB.First(&student, "id = ?", id).Error; err != nil {
		return false, "No Find Student", student
	}

	return true, "Find Student", student
}

// DeleteStudent deletes a student from the database by their ID.
// It takes the ID of the student as input and returns a tuple containing a boolean value indicating the success or failure of the operation,
// a string message describing the result, and the deleted Student object (if available).
//
// Parameters:
//
//	id (uint): The ID of the student to be deleted.
//
// Returns:
//
//	success (bool): A boolean value indicating the success (true) or failure (false) of the delete operation.
//	message (string): A string message providing information about the result of the delete operation.
//	student (models.Student): The deleted Student object, or an empty Student object if not found.
//
// Example:
//
//	success, message, deletedStudent := DeleteStudent(1)
//	if success {
//	    fmt.Println("Student deleted:", deletedStudent)
//	} else {
//	    fmt.Println("Error:", message)
//	}
func DeleteStudent(id uint) (bool, string, models.Student) {

	var student models.Student
	if err := config.DB.Delete(&student, "id = ?", id).Error; err != nil {
		return false, "No Find Student", student
	}

	return true, "Delete Student", student
}

// UpdateStudent updates the information of a student in the database.
// It takes the ID of the student and the new student data as input and returns a tuple containing a boolean value indicating the success or failure of the operation,
// a string message describing the result, and the updated Student object (if available).
//
// Parameters:
//
//	id (uint): The ID of the student to be updated.
//	new_student (schemas.StudentUpdateSchema): The new student data to be updated.
//
// Returns:
//
//	success (bool): A boolean value indicating the success (true) or failure (false) of the update operation.
//	message (string): A string message providing information about the result of the update operation.
//	student (models.Student): The updated Student object, or an empty Student object if not found.
//
// Example:
//
//	newStudent := schemas.StudentUpdateSchema{
//	    Name: "John",
//	    LastName: "Doe",
//	    Email: "john.doe@example.com",
//	    AccountNumber: "12345",
//	    Semester: 2,
//	    IdUniversityProgram: 1,
//	}
//	success, message, updatedStudent := UpdateStudent(1, newStudent)
//	if success {
//	    fmt.Println("Student updated:", updatedStudent)
//	} else {
//	    fmt.Println("Error:", message)
//	}
func UpdateStudent(id uint, new_student schemas.StudentUpdateSchema) (bool, string, models.Student) {

	var student models.Student
	if err := config.DB.First(&student, "id = ?", id).Error; err != nil {
		return false, "No Find Student", student
	} else {
		student.Name = tools.CopyField(new_student.Name, student.Name, "")
		student.LastName = tools.CopyField(new_student.LastName, student.LastName, "")
		student.Email = tools.CopyField(new_student.Email, student.Email, "")
		student.AccountNumber = tools.CopyField(new_student.AccountNumber, student.AccountNumber, "")
		student.Semester = tools.CopyField(new_student.Semester, student.Semester, 0)
		student.IdUniversityProgram = tools.CopyField(new_student.IdUniversityProgram, student.IdUniversityProgram, 0)

		config.DB.Save(&student)
	}

	return true, "Update Student Successful", student
}

// FindStudentByAccountNumber finds a student by their account number in the database.
// It takes the account number as input and returns a tuple containing a boolean value indicating the success or failure of the operation,
// a string message describing the result, and the found Student object (if available).
//
// Parameters:
//
//	accountNumber (string): The account number of the student to be found.
//
// Returns:
//
//	success (bool): A boolean value indicating the success (true) or failure (false) of the find operation.
//	message (string): A string message providing information about the result of the find operation.
//	student (models.Student): The found Student object, or an empty Student object if not found.
//
// Example:
//
//	success, message, foundStudent := FindStudentByAccountNumber("12345")
//	if success {
//	    fmt.Println("Student found:", foundStudent)
//	} else {
//	    fmt.Println("Error:", message)
//	}
func FindStudentByAccountNumber(accountNumber string) (bool, string, models.Student) {

	var student models.Student
	if err := config.DB.First(&student, "account_number = ?", accountNumber).Error; err != nil {
		return false, "No Find Student", student
	}

	return true, "Find Student", student
}
