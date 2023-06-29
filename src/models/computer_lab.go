package models

import "gorm.io/gorm"

type ComputerLab struct {
	gorm.Model
	// Metadata

	//Data
	Name string

	Description string
}
