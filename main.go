package main

import (
	"yorch-devs/bookstore-golang-postgres/dbutils"
	"yorch-devs/bookstore-golang-postgres/routes"

	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()
	dbutils.ConnectToDB()
	// seed.RunSeeds()

	router := gin.Default()
	routes.MountRoutes(router)
	router.Run("0.0.0.0:8000")
}
