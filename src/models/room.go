package models

import "github.com/jinzhu/gorm"

type Room struct {
	gorm.Model
	// Metadata
	idRoom string
	idUser string

	//Data
	name string
}
