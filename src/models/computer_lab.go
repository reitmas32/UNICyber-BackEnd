package models

import "github.com/jinzhu/gorm"

type ComputerLab struct {
	gorm.Model
	// Metadata
	IdComputerLab string `gorm:"primaryKey"`

	//Data
	Name string

	Description string
}
