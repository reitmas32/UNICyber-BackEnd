package models

import "github.com/jinzhu/gorm"

type LinkAccountCode struct {
	gorm.Model
	// Metadata
	UserName      string `json:"user_name"`
	Code          string `json:"code" binding:"len=6"`
	IdComputerLab uint   `json:"id_computer_lab"`
}
