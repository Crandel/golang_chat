package models

import (
	"app/utils/db"
	"log"
	"time"
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
func (m *Message) SaveMessage() (uint, error) {
	dbase := db.DB
	m.CreatedAt = time.Now()
	err := dbase.Save(m).Error
	if err != nil {
		return 0, err
	}
	return m.ID, nil
}

// GetMessages ...
func GetMessages(m *[]Message) error {
	dbase := db.DB
	now := time.Now()
	year, month, day := now.Date()
	today := time.Date(year, month, day, 0, 0, 0, 0, now.Location())
	log.Println(today)
	return dbase.Preload("User").Where("created_at >= ?", today).Order("id asc").Find(m).Error
}

// CreateUser ...
func (user *User) CreateUser() error {
	dbase := db.DB
	return dbase.Create(user).Error
}

// GetMessage - return message using id
func (m *Message) GetMessage() error {
	dbase := db.DB
	return dbase.First(m, m.ID).Error
}

// DeleteMessage delete message, must be id!
func (m *Message) DeleteMessage() error {
	dbase := db.DB
	return dbase.Delete(m).Error
}
