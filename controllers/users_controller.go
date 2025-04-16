package controllers

import (
	"net/http"
	"yorch-devs/bookstore-golang-postgres/dbutils"
	"yorch-devs/bookstore-golang-postgres/models"

	"github.com/gin-gonic/gin"
)

type UserSingleRecord struct {
	User         models.UserInfo `json:"user,omitempty"`
	Error        string          `json:"error,omitempty"`
	RowsAffected int64           `json:"rows_affected,omitempty"`
}

func GetUser(c *gin.Context) {
	var user models.User
	var userSR UserSingleRecord
	id := c.Param("id")

	result := dbutils.Db.First(&user, "id = ?", id)

	if result.Error != nil {
		userSR.Error = result.Error.Error()
		c.JSON(http.StatusBadRequest, gin.H{"error": userSR.Error})
		return
	}

	userSR.User.Id = user.Id
	userSR.User.Login = user.Login
	userSR.User.CreatedAt = user.CreatedAt
	userSR.User.UpdatedAt = user.UpdatedAt
	userSR.RowsAffected = result.RowsAffected
	c.JSON(http.StatusOK, userSR)
}

func SignUp(c *gin.Context) {
	var user models.User
	var userSR UserSingleRecord

	if err := c.ShouldBindJSON(&user); err != nil {
		userSR.Error = err.Error()
		c.JSON(http.StatusBadRequest, gin.H{"error": userSR.Error})
		return
	}

	result := dbutils.Db.Create(&user)

	if result.Error != nil {
		userSR.Error = result.Error.Error()
		c.JSON(http.StatusBadRequest, gin.H{"error": userSR.Error})
		return
	}

	userSR.User.Id = user.Id
	userSR.User.Login = user.Login
	userSR.User.CreatedAt = user.CreatedAt
	userSR.User.UpdatedAt = user.UpdatedAt
	userSR.RowsAffected = result.RowsAffected
	c.JSON(http.StatusOK, userSR)
}

func DeleteUser(c *gin.Context) {
	var user models.User
	var userSR UserSingleRecord
	id := c.Param("id")

	result := dbutils.Db.First(&user, "id = ?", id)

	if result.Error != nil {
		userSR.Error = result.Error.Error()
		c.JSON(http.StatusNotFound, gin.H{"error": userSR.Error})
		return
	}

	result = dbutils.Db.Delete(&user)

	if result.Error != nil {
		userSR.Error = result.Error.Error()
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": userSR.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "The user was deleted successfully"})
}
