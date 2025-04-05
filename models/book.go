package models

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	Id        uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid();unique" json:"id"`
	Title     string    `gorm:"not null;size:256;index;unique" json:"title" binding:"required"`
	Author    string    `gorm:"not null;size:256;index;" json:"author" binding:"required"`
	CreatedAt time.Time `gorm:"not null;column:created_at;default:current_timestamp;index" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null;column:updated_at;default:current_timestamp;index" json:"updated_at"`
}

func (Book) TableName() string {
	return "books"
}
