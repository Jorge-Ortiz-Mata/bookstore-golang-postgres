package routes

import (
	"yorch-devs/bookstore-golang-postgres/controllers"

	"github.com/gin-gonic/gin"
)

func MountRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/books", controllers.GetBooks)
		v1.GET("/books/:id", controllers.GetBook)
		v1.PATCH("/books/:id", controllers.UpdateBook)
		v1.POST("/books", controllers.CreateBook)
		v1.DELETE("/books/:id", controllers.DeleteBook)

		v1.GET("/users/:id", controllers.GetUser)
		v1.POST("/users/signup", controllers.SignUp)
		v1.DELETE("/users/:id", controllers.DeleteUser)

		v1.POST("/login", controllers.Login)
	}
}
