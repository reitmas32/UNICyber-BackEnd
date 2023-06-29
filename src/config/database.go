package config

import (
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/models"
	//"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func SetupDB() {
	//DB, err = gorm.Open("mysql", "rafael:tere@tcp(127.0.0.1:4306)/sisec_test?charset=utf8mb4&parseTime=True&loc=Local")
	//DB, err = gorm.Open("sqlite3", "app/my_database.db")
	dsn := "my_database.db"
	DB, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		print(err)
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
	DB.AutoMigrate(&models.UniversityProgram{})

	//Operations
	DB.AutoMigrate(&models.Loan{})

}
