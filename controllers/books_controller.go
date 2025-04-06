package controllers

import (
	"net/http"
	"yorch-devs/bookstore-golang-postgres/dbutils"
	"yorch-devs/bookstore-golang-postgres/models"

	"github.com/gin-gonic/gin"
)

type BookSingleRecord struct {
	Book         *models.Book `json:"book,omitempty"`
	Error        string       `json:"error,omitempty"`
	RowsAffected int64        `json:"rows_affected,omitempty"`
}

type BooksMultipleRecords struct {
	Books        *[]models.Book `json:"books,omitempty"`
	Error        string         `json:"error,omitempty"`
	RowsAffected int64          `json:"rows_affected,omitempty"`
}

func GetBooks(c *gin.Context) {
	var books []models.Book
	var booksMR BooksMultipleRecords

	result := dbutils.Db.Find(&books)

	if result.Error != nil {
		booksMR.Error = result.Error.Error()
		c.JSON(http.StatusBadRequest, booksMR)
		return
	}

	booksMR.Books = &books
	booksMR.RowsAffected = result.RowsAffected
	c.JSON(http.StatusOK, booksMR)
}

func GetBook(c *gin.Context) {
	var book models.Book
	var bookSR BookSingleRecord
	id := c.Param("id")

	result := dbutils.Db.First(&book, "id = ?", id)

	if result.Error != nil {
		bookSR.Error = result.Error.Error()
		c.JSON(http.StatusBadRequest, gin.H{"error": bookSR.Error})
		return
	}

	bookSR.Book = &book
	bookSR.RowsAffected = result.RowsAffected

	c.JSON(http.StatusOK, bookSR)
}

func CreateBook(c *gin.Context) {
	var book models.Book
	var bookSR BookSingleRecord

	if err := c.ShouldBindJSON(&book); err != nil {
		bookSR.Error = err.Error()
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": bookSR.Error})
		return
	}

	result := dbutils.Db.Create(&book)

	if result.Error != nil {
		bookSR.Error = result.Error.Error()
		c.JSON(http.StatusBadRequest, gin.H{"error": bookSR.Error})
		return
	}

	bookSR.RowsAffected = result.RowsAffected
	bookSR.Book = &book
	c.JSON(http.StatusOK, bookSR)
}

func UpdateBook(c *gin.Context) {
	var book models.Book
	var data map[string]any
	var bookSR BookSingleRecord
	id := c.Param("id")

	result := dbutils.Db.First(&book, "id = ?", id)

	if result.Error != nil {
		bookSR.Error = result.Error.Error()
		c.JSON(http.StatusBadRequest, gin.H{"error": bookSR.Error})
		return
	}

	if err := c.ShouldBindJSON(&data); err != nil {
		bookSR.Error = err.Error()
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": bookSR.Error})
		return
	}

	result = dbutils.Db.Model(&book).Updates(data)

	if result.Error != nil {
		bookSR.Error = result.Error.Error()
		c.JSON(http.StatusInternalServerError, gin.H{"error": bookSR.Error})
		return
	}

	bookSR.Book = &book
	bookSR.RowsAffected = result.RowsAffected

	c.JSON(http.StatusOK, bookSR)
}

func DeleteBook(c *gin.Context) {
	var book models.Book
	var bookSR BookSingleRecord
	id := c.Param("id")

	result := dbutils.Db.First(&book, "id = ?", id)

	if result.Error != nil {
		bookSR.Error = result.Error.Error()
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": bookSR.Error})
		return
	}

	result = dbutils.Db.Delete(&book)

	if result.Error != nil {
		bookSR.Error = result.Error.Error()
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": bookSR.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "The book has been deleted successfully"})
}
