package models

import (
	"app/utils/db"
	"log"
)

// Automigrate ...
func Automigrate() {
	dbase := db.DB
	// Automigration of tables
	dbase.AutoMigrate(&User{}, &Message{})
}

// GetUserByLoginPass ...
func (user *User) GetUserByLoginPass(login string, pass string) bool {
	dbase := db.DB
	err := dbase.Where(&User{Login: login, Password: pass}).First(user).RecordNotFound()
	return err
}

// GetUserByID ...
func (user *User) GetUserByID(id uint) {
	dbase := db.DB
	dbase.First(user, id)
}

// SaveMessage - save single message
func (user *User) SaveMessage(m string) {
	dbase := db.DB
	log.Printf("%#v\n", user)
	message := Message{User: *user, Message: m}
	dbase.Save(message)
	messages := user.Messages
	messages = append(messages, message)
	user.Messages = messages
	log.Printf("Message: %#v\n\n User: %#v\n", message, user)
	dbase.Save(user)
}

// GetMessages ...
func GetMessages(m *[]Message) {
	dbase := db.DB
	dbase.Preload("User").Find(m)
}

// CreateUser ...
func (user *User) CreateUser() {
	dbase := db.DB
	dbase.Create(user)
}
