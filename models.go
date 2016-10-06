package main

import "time"

// User model from database
type User struct {
	ID       int
	Email    string
	Login    string
	Password string
}

// Message from database
type Message struct {
	User    User
	Message string
	Time    time.Time
}
