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
	host := os.Getenv("POSTGRES_HOST")
	username := os.Getenv("POSTGRES_USERNAME")
	password := os.Getenv("POSTGRES_PASSWORD")
	database := os.Getenv("POSTGRES_DATABASE")

	// // Dev
	// dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=5432 sslmode=disable", host, username, password, database)

	// // Prod
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v", host, username, password, database)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println("The database was connected successfully")
	Db = db
}
