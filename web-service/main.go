package main

import (
	"yorch-devs/bookstore-golang-postgres/dbutils"
	"yorch-devs/bookstore-golang-postgres/routes"
)

func main() {
	dbutils.ConnectToDB()
	dbutils.ApplyMigrations(dbutils.Db)
	routes.MountRoutes()
}
