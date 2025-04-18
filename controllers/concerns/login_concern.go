package concerns

import (
	"yorch-devs/bookstore-golang-postgres/controllers/controller_structs"
	"yorch-devs/bookstore-golang-postgres/dbutils"
	"yorch-devs/bookstore-golang-postgres/models"

	"golang.org/x/crypto/bcrypt"
)

func ValidateUsername(username string) controller_structs.UsernameValidationResult {
	var usernameVR controller_structs.UsernameValidationResult

	result := dbutils.Db.First(&usernameVR.User, "username = ?", username)
	usernameVR.Result = result

	return usernameVR
}

func ValidatePassword(password string, user *models.User) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
