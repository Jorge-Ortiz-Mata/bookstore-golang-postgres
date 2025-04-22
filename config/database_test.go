package config

import (
	"fmt"
	"testing"
)

func TestDevelopmentDatabase(t *testing.T) {
	host := "localhost"
	port := "5432"
	username := "jorge_test"
	password := "jorge_test"
	database := "bookstore_golang_postgres_development"

	var dsn string = DevelopmentDatabase()
	var databaseConnection string = fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable", host, username, password, database, port)

	if dsn != databaseConnection {
		t.Errorf("The development database dsn is invalid.\nexpected: %v\ngot: %v", databaseConnection, dsn)
	}
}

func TestTestingDatabase(t *testing.T) {
	host := "localhost"
	port := "5432"
	username := "jorge_test"
	password := "jorge_test"
	database := "bookstore_golang_postgres_test"

	var dsn string = TestingDatabase()
	var databaseConnection string = fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable", host, username, password, database, port)

	if dsn != databaseConnection {
		t.Errorf("The testing database dsn is invalid.\nexpected: %v\ngot: %v", databaseConnection, dsn)
	}
}

func TestProductionDatabase(t *testing.T) {
	host := "localhost"
	username := "jorge_test"
	password := "jorge_test"
	database := "bookstore_golang_postgres_production"

	var dsn string = ProductionDatabase()
	var databaseConnection string = fmt.Sprintf("host=%v user=%v password=%v dbname=%v", host, username, password, database)

	if dsn != databaseConnection {
		t.Errorf("The production database dsn is invalid.\nexpected: %v\n got: %v\n", databaseConnection, dsn)
	}
}
