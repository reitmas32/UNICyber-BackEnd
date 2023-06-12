package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Loan struct {
	gorm.Model

	IdStudent  uint
	IdComputer uint

	SesionStart time.Time
	SesionEnd   time.Time
}
