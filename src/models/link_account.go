package models

import "github.com/jinzhu/gorm"

type LinkAccount struct {
	gorm.Model
	// Metadata
	UserName      string `json:"user_name"`
	IdComputerLab uint   `json:"id_computer_lab"`
}
