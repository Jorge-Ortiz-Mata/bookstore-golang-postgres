package controllers

import (
	"fmt"
	"net/http"
	"yorch-devs/bookstore-golang-postgres/models"

	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context) {
	var book models.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(book)
	c.JSON(http.StatusOK, book)
}
