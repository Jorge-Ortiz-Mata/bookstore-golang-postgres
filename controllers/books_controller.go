package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"yorch-devs/bookstore-golang-postgres/dbutils"
	"yorch-devs/bookstore-golang-postgres/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

const DefaultLimit = 10

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

func setAfterIdParam(c *gin.Context) (string, error) {
	afterId, afterIdExists := c.GetQuery("after_id")

	if !afterIdExists {
		return "", nil
	}

	if err := uuid.Validate(afterId); err != nil {
		return "", errors.New("invalid after uuid")
	}

	return afterId, nil
}

func GetBooks(c *gin.Context) {
	var books []models.Book
	var booksMR BooksMultipleRecords
	var result *gorm.DB

	limit, err := setLimitParam(c)

	if err != nil {
		booksMR.Error = err.Error()
		c.JSON(http.StatusBadRequest, gin.H{"error": booksMR.Error})
		return
	}

	after_id, err := setAfterIdParam(c)

	if err != nil {
		booksMR.Error = err.Error()
		c.JSON(http.StatusBadRequest, gin.H{"error": booksMR.Error})
		return
	}

	if len(after_id) > 0 {
		fmt.Println(after_id)
		result = dbutils.Db.Order("id asc").Where("id > ?", after_id).Limit(limit).Find(&books)
	} else {
		result = dbutils.Db.Order("id asc").Limit(limit).Find(&books)
	}

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

func setLimitParam(c *gin.Context) (int, error) {
	limitStr, limitStrExists := c.GetQuery("limit")

	if limitStrExists {
		limit, err := strconv.Atoi(limitStr)

		if err != nil {
			return 0, err
		}

		if limit <= 0 || limit > 10 {
			return DefaultLimit, nil
		}

		return limit, nil
	}

	return DefaultLimit, nil
}
