package seed

import (
	"log"
	"yorch-devs/bookstore-golang-postgres/dbutils"
	"yorch-devs/bookstore-golang-postgres/models"

	"github.com/jaswdr/faker"
)

var bookTitles []string = []string{
	"The Silent Forest", "Echoes of Tomorrow", "Whispers in the Dark", "Beneath Crimson Skies", "The Forgotten Path",
	"Shadows of the Mind", "A World Apart", "The Last Ember", "Tales from the North", "Beyond the Horizon",
	"Fragments of Truth", "The Painted Veil", "Winds of Fate", "The Midnight Tower", "Letters from Nowhere",
	"The Glass Kingdom", "Steps into Shadow", "The Cursed Crown", "The Final Hour", "Bridge to the Stars",
}

func SeedBooks() {
	log.Println("Seeding books...")

	f := faker.New()

	for i := range len(bookTitles) {
		var book models.Book
		book.Title = bookTitles[i]
		book.Author = f.Person().Name()
		dbutils.Db.Create(&book)
	}

	log.Println("Books created")
}
