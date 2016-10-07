package main

import (
	"time"

	"github.com/jinzhu/gorm"
)

// User model from database
type User struct {
	gorm.Model
	Email    string `gorm:"size:50"`
	Login    string `gorm:"size:50"`
	Password string `gorm:"size:255"`
	Messages []Message
}

// Message from database
type Message struct {
	ID        uint
	UserID    uint
	User      User
	CreatedAt time.Time
	Message   string
}
