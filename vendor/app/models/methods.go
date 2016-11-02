package models

import "app/utils/db"

// Automigrate ...
func Automigrate() {
	dbase := db.DB
	// Automigration of tables
	dbase.AutoMigrate(&User{}, &Message{})
}

// GetUserByEmailPass ...
func GetUserByEmailPass(user *User) bool {
	dbase := db.DB
	err := dbase.Where(user).First(user).RecordNotFound()
	return err
}

// GetMessages ...
func GetMessages(m *[]Message) {
	dbase := db.DB
	dbase.Preload("User").Find(m)
}
