package models

import (
	"time"

	"github.com/google/uuid"
)

type UserInfo struct {
	Id        uuid.UUID `json:"id,omitempty"`
	Login     string    `json:"login"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type User struct {
	Id        uuid.UUID `gorm:"default:gen_random_uuid()" json:"id,omitempty"`
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
