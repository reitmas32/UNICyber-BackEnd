package models

import "github.com/jinzhu/gorm"

type Room struct {
	gorm.Model
	// Metadata
	IdComputerLab uint

	//Data
	Name string

	Index int
}
