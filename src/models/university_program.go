package models

import "gorm.io/gorm"

type UniversityProgram struct {
	gorm.Model

	//Data
	Name string

	//Image
	//Image string
}
