package dbutils

import (
	"log"
	"yorch-devs/bookstore-golang-postgres/models"

	"gorm.io/gorm"
)

type Tabler interface {
	TableName() string
}

func ApplyMigrations(Db *gorm.DB) {
	err := Db.AutoMigrate(&models.Book{})

	if err != nil {
		log.Fatalf("Something went wrong while running book migration")
	}
}
