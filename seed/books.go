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
	"The Lost Heir", "Rise of the Phoenix", "Dance of Ashes", "The Ivory Gate", "Memories of the Deep",
	"Guardians of the Flame", "A Time to Wander", "Voices of the Wind", "The Moonlit Vale", "Storm over Eden",
	"The Wandering Flame", "Thorns and Petals", "A Kingdom Reborn", "The Secret Harvest", "Tide of Echoes",
	"The Winter Pact", "Chronicles of Dust", "The Wishing Stone", "Legacy of Fire", "The Labyrinth Code",
	"Portraits of Silence", "Garden of Shadows", "The Hollow Keep", "Realm of Whispers", "Bonds of Smoke",
	"Feathers and Steel", "Beneath the Gilded Sky", "Tears of the Mountain", "The Sapphire Oath", "Blood and Bloom",
	"The Long Nightfall", "Lanterns in the Mist", "The Crimson Pact", "A Spark Among Ruins", "The Forgotten Waltz",
	"Court of Ash and Snow", "The Clockmaker's Prophecy", "Frost and Flame", "The Raven's Map", "Crownless Kings",
	"The Whispering Flame", "The Mirror Gate", "Dust in the Rain", "The Gilded Shadow", "Roots of the Fallen",
	"Driftwood Memories", "The Burning Script", "Salt in the Wound", "The Inkwell Pact", "Harvest of Echoes",
	"Twilight Rebellion", "A Song for the Lost", "The Obsidian Seal", "Scars of the Moon", "The Hollow Star",
	"Fire Among Thorns", "The Serpent's Veil", "Moonrise Over Glass", "Threads of Destiny", "The Paper Kingdom",
	"Echoes of the Sea", "The Silver Testament", "The Veil Between", "Dawn over Emberfall", "Vows in the Dark",
	"The Bone Orchard", "The Forgotten Canvas", "The Black Feather", "A Flame at Dusk", "Song of the Hollow",
	"The Booksmith's Curse", "The King's Echo", "Dust and Silence", "The Unseen Garden", "The Last Archivist",
	"The Lantern War", "Beneath the Thorn Tree", "Tides of Ash", "The Crescent War", "Notes from the Shadows",
	"Castle of Winter", "Embers of the Past", "The Inked Vow", "Oath to the Sky", "The Frozen Lute",
}

func SeedBooks() {
	log.Println("Seeding books...")

	f := faker.New()

	for i := 0; i < 100; i++ {
		var book models.Book
		book.Title = bookTitles[i]
		book.Author = f.Person().Name()
		dbutils.Db.Create(&book)
	}

	log.Println("Books created")
}
