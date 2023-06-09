package models

import "github.com/jinzhu/gorm"

type Computer struct {
	gorm.Model
	// Metadata
	IdRoom uint

	// UI
	Pos_x float32
	Pos_y float32

	// Data
	Name    string
	State   string
	IdState uint
	Message string
	Type    string
}
