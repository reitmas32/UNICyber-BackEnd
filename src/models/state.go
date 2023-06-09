package models

import "github.com/jinzhu/gorm"

type State struct {
	gorm.Model

	//Data
	Name string
}
