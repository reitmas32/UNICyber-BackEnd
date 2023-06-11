package models

import "github.com/jinzhu/gorm"

type UniversityProgram struct {
	gorm.Model

	//Data
	Name string

	//Image
	//Image string
}
