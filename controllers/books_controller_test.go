package controllers_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"yorch-devs/bookstore-golang-postgres/dbutils"
	"yorch-devs/bookstore-golang-postgres/models"
	"yorch-devs/bookstore-golang-postgres/routes"
	"yorch-devs/bookstore-golang-postgres/seed"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var router *gin.Engine = gin.Default()

func init() {
	os.Setenv("ENV", "testing")
	dbutils.ConnectToDB()
	routes.MountRoutes(router)
}

func CleanUpBooks() {
	dbutils.Db.Exec("DELETE FROM books")
}

type BooksResponse struct {
	Books        []models.Book `json:"books"`
	RowsAffected int           `json:"rows_affected"`
}

func TestGetBooks(t *testing.T) {
	seed.SeedBooks()

	request, _ := http.NewRequest("GET", "/api/v1/books", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
	responseData, _ := io.ReadAll(response.Body)

	var data BooksResponse
	json.Unmarshal(responseData, &data)

	assert.Equal(t, len(data.Books), 10, "The books array length must be 10")
	assert.Equal(t, data.RowsAffected, 10, "The rows affected value must be 10")

	CleanUpBooks()
}
