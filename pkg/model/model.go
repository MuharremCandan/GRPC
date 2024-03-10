package model

import (
	"time"
)

// User struct represents a user with basic information
type User struct {
	ID        uint32    `gorm:"primaryKey" json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
