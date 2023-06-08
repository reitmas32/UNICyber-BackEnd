package models

import "github.com/jinzhu/gorm"

type ComputerLab struct {
	gorm.Model
	// Metadata

	//Data
	Name string

	Description string
}
