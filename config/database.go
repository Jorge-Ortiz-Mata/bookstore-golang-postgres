package config

import (
	"fmt"
	"os"
)

func DevelopmentDatabase() string {
	host := os.Getenv("PGHOST")
	port := os.Getenv("PGPORT")
	username := os.Getenv("PGUSER")
	password := os.Getenv("PGPASSWORD")
	database := os.Getenv("DEVELOPMENT_DATABASE")

	return fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable", host, username, password, database, port)
}

func TestingDatabase() string {
	host := os.Getenv("PGHOST")
	port := os.Getenv("PGPORT")
	username := os.Getenv("PGUSER")
	password := os.Getenv("PGPASSWORD")
	database := os.Getenv("TEST_DATABASE")

	return fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable", host, username, password, database, port)
}

func ProdutionDatbase() string {
	host := os.Getenv("PGHOST")
	username := os.Getenv("PGUSER")
	password := os.Getenv("PGPASSWORD")
	database := os.Getenv("PRODUCTION_DATABASE")

	return fmt.Sprintf("host=%v user=%v password=%v dbname=%v", host, username, password, database)
}
