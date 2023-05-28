package models

import "github.com/jinzhu/gorm"

type Student struct {
	gorm.Model
	idUser   string
	Password string
	UserName string
}
