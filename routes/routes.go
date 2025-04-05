package routes

import (
	"yorch-devs/bookstore-golang-postgres/controllers"

	"github.com/gin-gonic/gin"
)

func MountRoutes() {
	router := gin.Default()
	router.POST("/books", controllers.CreateBook)
	router.Run("localhost:8000")
}
