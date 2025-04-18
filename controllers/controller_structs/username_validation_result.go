package controller_structs

import (
	"yorch-devs/bookstore-golang-postgres/models"

	"gorm.io/gorm"
)

type UsernameValidationResult struct {
	User   *models.User
	Result *gorm.DB
}
