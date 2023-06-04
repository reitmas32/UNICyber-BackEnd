package models

import "github.com/jinzhu/gorm"

type Computer struct {
	gorm.Model
	// Metadata
	IdComputer string `gorm:"primaryKey"`
	IdRoom     string

	// UI
	Pos_x float32
	Pos_y float32

	// Data
	Name    string
	State   string
	Message string
	Type    string
}