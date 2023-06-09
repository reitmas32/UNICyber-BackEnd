package config

import (
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/models"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB
var err error

func SetupDB() {
	DB, err = gorm.Open("sqlite3", "my_database.db")
	if err != nil {
		panic(err)
	}
}

func MigrateDB() {
	// Genera las tablas en la base de datos
	DB.AutoMigrate(&models.Computer{})
	DB.AutoMigrate(&models.Room{})
	DB.AutoMigrate(&models.Student{})
	DB.AutoMigrate(&models.ComputerLab{})

	//Tables Intermedas
	DB.AutoMigrate(&models.LinkAccount{})

	//Tools
	DB.AutoMigrate(&models.LinkAccountCode{})

	//State
	DB.AutoMigrate(&models.State{})

}
