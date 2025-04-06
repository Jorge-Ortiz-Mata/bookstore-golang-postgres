package routes

import (
	"yorch-devs/bookstore-golang-postgres/controllers"

	"github.com/gin-gonic/gin"
)

func MountRoutes() {
	router := gin.Default()
	router.GET("/books", controllers.GetBooks)
	router.GET("/books/:id", controllers.GetBook)
	router.PATCH("/books/:id", controllers.UpdateBook)
	router.POST("/books", controllers.CreateBook)
	router.Run("localhost:8000")
}
