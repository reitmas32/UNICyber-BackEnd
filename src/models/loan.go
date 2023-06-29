package models

import (
	"time"

	"gorm.io/gorm"
)

type Loan struct {
	gorm.Model

	IdStudent  uint
	IdComputer uint

	SesionStart time.Time
	SesionEnd   time.Time
}
