package models

import (
	"time"

	"github.com/google/uuid"
)

type UserInfo struct {
	Id        uuid.UUID `json:"id,omitempty"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type User struct {
	Id        uuid.UUID `gorm:"default:gen_random_uuid()" json:"id,omitempty"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
