package models

import "app/utils/db"

// Automigrate ...
func Automigrate() {
	dbase := db.DB
	// Automigration of tables
	dbase.AutoMigrate(&User{}, &Message{})
}

// GetUserByLoginPass ...
func GetUserByLoginPass(login string, pass string, user *User) bool {
	dbase := db.DB
	err := dbase.Where(&User{Login: login, Password: pass}).First(user).RecordNotFound()
	return err
}

// GetMessages ...
func GetMessages(m *[]Message) {
	dbase := db.DB
	dbase.Preload("User").Find(m)
}

// CreateUser ...
func CreateUser(user *User) {
	dbase := db.DB
	dbase.Create(user)
}
