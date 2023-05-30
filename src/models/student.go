package models

import "github.com/jinzhu/gorm"

type Student struct {
	gorm.Model
	//Metadata
	IdStudent string `gorm:"primaryKey"`

	//Info Personal
	Name     string
	LastName string
	Email    string

	//Info Academic
	UniversityProgram string
	AccountNumber     string
	Semester          uint8
}
