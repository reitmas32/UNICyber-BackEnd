package models

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	// Metadata
	IdComputerLab uint

	//Data
	Name string

	Index int
}
