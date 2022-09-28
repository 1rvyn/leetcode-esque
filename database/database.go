package database

import (
	"fiberWebApi/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

type Dbinstance struct {
	Db *gorm.DB
}

var Database Dbinstance

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect to the database \n", err.Error())
		os.Exit(2)

	}

	log.Println("there was a successful connection to the Database")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations")
	// TODO: add migrations

	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})

	Database = Dbinstance{Db: db}

}
