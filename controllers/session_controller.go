package controllers

import (
	"net/http"
	"yorch-devs/bookstore-golang-postgres/controllers/concerns"
	"yorch-devs/bookstore-golang-postgres/controllers/controller_structs"
	"yorch-devs/bookstore-golang-postgres/utils"

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

	token, err := utils.GenerateAuthToken(validUsername.User.Id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong when generating auth token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"auth_token": token})
}
