package models

import "gorm.io/gorm"

type State struct {
	gorm.Model

	//Data
	Name string

	//Image
	Image string
}
