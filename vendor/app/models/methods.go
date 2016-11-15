package models

import (
	"app/utils/db"
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
	return dbase.Where(&User{Login: login, Password: pass}).First(user).RecordNotFound()
}

// GetUserByID ...
func GetUserByID(id uint) (*User, error) {
	dbase := db.DB
	user := &User{}
	err := dbase.First(user, id).Error
	return user, err
}

// SaveMessage - save single message
func (user *User) SaveMessage(m string) (uint, error) {
	dbase := db.DB
	message := Message{UserID: user.ID, Message: m}
	err := dbase.Save(&message).Error
	if err != nil {
		return 0, err
	}
	return message.ID, nil
}

// GetMessages ...
func GetMessages(m *[]Message) error {
	dbase := db.DB
	return dbase.Preload("User").Find(m).Error
}

// CreateUser ...
func (user *User) CreateUser() error {
	dbase := db.DB
	return dbase.Create(user).Error
}

// GetMessage - return message using id
func (m *Message) GetMessage(id uint) error {
	dbase := db.DB
	return dbase.First(m, id).Error
}
