package seed

import (
	"log"
	"yorch-devs/bookstore-golang-postgres/dbutils"
	"yorch-devs/bookstore-golang-postgres/models"

	"github.com/jaswdr/faker"
)

func SeedBooks() {
	log.Println("Seeding books...")

	f := faker.New()

	for range 380 {
		var book models.Book
		book.Title = f.Lorem().Sentence(3)
		book.Author = f.Person().Name()
		dbutils.Db.Create(&book)
	}

	log.Println("Books created")
}
