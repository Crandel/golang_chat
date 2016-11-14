package models

import "app/utils/db"

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
func (user *User) SaveMessage(m string) uint {
	dbase := db.DB
	message := Message{UserID: user.ID, Message: m}
	dbase.Save(&message)
	return message.ID
}

// GetMessages ...
func GetMessages(m *[]Message) {
	dbase := db.DB
	dbase.Preload("User").Find(m)
}

// CreateUser ...
func (user *User) CreateUser() error {
	dbase := db.DB
	return dbase.Create(user).Error
}
