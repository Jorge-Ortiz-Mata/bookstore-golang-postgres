package main

import (
	"yorch-devs/bookstore-golang-postgres/dbutils"
	"yorch-devs/bookstore-golang-postgres/routes"

	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()
	dbutils.ConnectToDB()
	dbutils.ApplyMigrations(dbutils.Db)
	routes.MountRoutes()
}
