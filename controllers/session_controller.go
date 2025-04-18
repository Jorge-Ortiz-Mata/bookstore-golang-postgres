package controllers

import (
	"net/http"
	"yorch-devs/bookstore-golang-postgres/controllers/concerns"
	"yorch-devs/bookstore-golang-postgres/controllers/controller_structs"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var login *controller_structs.Login

	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse JSON object"})
		return
	}

	validUsername := concerns.ValidateUsername(login.Username)

	if validUsername.Result.Error != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid username"})
		return
	}

	validPassword := concerns.ValidatePassword(login.Password, validUsername.User)

	if !validPassword {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User has successfully authenticated"})
}
