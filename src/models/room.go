package models

import "github.com/jinzhu/gorm"

type Room struct {
	gorm.Model
	// Metadata
	IdRoom string `gorm:"primaryKey"`
	IdComputerLab string

	//Data
	Name string

	Index int
}
