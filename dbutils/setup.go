package dbutils

import (
	"log"
	"os"
	"yorch-devs/bookstore-golang-postgres/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func ConnectToDB() {
	env := os.Getenv("ENV")
	var dsn string

	if env == "development" {
		dsn = config.DevelopmentDatabase()
	} else if env == "production" {
		dsn = config.ProductionDatabase()
	} else {
		dsn = config.TestingDatabase()
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println("The database was connected successfully")
	Db = db
}
