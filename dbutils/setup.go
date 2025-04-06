package dbutils

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func ConnectToDB() {
	username := os.Getenv("POSTGRES_USERNAME")
	password := os.Getenv("POSTGRES_PASSWORD")
	dsn := fmt.Sprintf("host=localhost user=%v password=%v dbname=jorge_test_db port=5432 sslmode=disable", username, password)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println("The database was connected successfully")
	Db = db
}
