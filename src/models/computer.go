package models

import "github.com/jinzhu/gorm"

type Computer struct {
	gorm.Model
	// Metadata
	idComputerUI string
	idRoom       string

	// UI
	pos_x int
	pos_y int

	// Data
	name         string
	state        string
	message      string
	typeComputer string
}
