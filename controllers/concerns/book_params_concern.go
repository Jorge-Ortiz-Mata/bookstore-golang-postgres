package concerns

import (
	"yorch-devs/bookstore-golang-postgres/dbutils"
	"yorch-devs/bookstore-golang-postgres/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const (
	DefaultLimit          int    = 10
	DefaultAttributeOrder string = "id"
	DefaultDirectionOrder string = "asc"
)

func DirectionWithOrder(orderBy, direction string, query *gorm.DB, book models.Book) *gorm.DB {
	switch orderBy {
	case "title":
		if direction == "asc" {
			return query.Where("title > ?", book.Title)
		}

		return query.Where("title < ?", book.Title)
	default:
		if direction == "asc" {
			return query.Where("author > ?", book.Author)
		}

		return query.Where("author < ?", book.Author)
	}
}

func OrderWithParams(afterId, orderBy, direction string, query *gorm.DB) *gorm.DB {
	var prevBook models.Book

	if afterId != "" && orderBy == "id" {
		if direction == "asc" {
			return query.Where("id > ?", afterId)
		}

		return query.Where("id < ?", afterId)
	}

	if afterId != "" && orderBy != "id" {
		dbutils.Db.First(&prevBook, "id = ?", afterId)

		return DirectionWithOrder(orderBy, direction, query, prevBook)
	}

	return query
}

func BooksDbQuery(c *gin.Context) ([]models.Book, int64, error) {
	var books []models.Book
	OrderExpression, orderBy, direction := OrderParam(c)
	query := dbutils.Db.Order(OrderExpression)

	if query.Error != nil {
		return nil, 0, query.Error
	}

	afterId, err := AfterIdParam(c)
	if err != nil {
		return nil, 0, err
	}

	limit, err := LimitParam(c)
	if err != nil {
		return nil, 0, err
	}

	query = OrderWithParams(afterId, orderBy, direction, query)

	if query.Error != nil {
		return nil, 0, query.Error
	}

	query = query.Limit(limit).Find(&books)

	if query.Error != nil {
		return nil, 0, query.Error
	}

	return books, query.RowsAffected, nil
}
