package controllers

import (
	"net/http"
	"os"
	"time"
	"yorch-devs/bookstore-golang-postgres/controllers/concerns"
	"yorch-devs/bookstore-golang-postgres/controllers/controller_structs"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
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

	token, err := generateAuthToken(validUsername.User.Id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong when generating auth token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"auth_token": token})
}

func generateAuthToken(userId uuid.UUID) (string, error) {
	var jwtSecret = []byte(os.Getenv("SECRET_KEY"))

	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(time.Second * 30).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecret)

	if err != nil {
		return "", err
	}

	return signedToken, nil
}
