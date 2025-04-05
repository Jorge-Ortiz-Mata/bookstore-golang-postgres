package models

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	Id        uuid.UUID `json:"id"`
	Title     string    `json:"title" binding:"required"`
	Author    string    `json:"author" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
