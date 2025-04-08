package routes

import (
	"yorch-devs/bookstore-golang-postgres/controllers"

	"github.com/gin-gonic/gin"
)

func MountRoutes() {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/books", controllers.GetBooks)
		v1.GET("/books/:id", controllers.GetBook)
		v1.PATCH("/books/:id", controllers.UpdateBook)
		v1.POST("/books", controllers.CreateBook)
		v1.DELETE("/books/:id", controllers.DeleteBook)
	}

	router.Run("0.0.0.0:8000")
}
