package db

import (
	m "app/models"
	"log"

	"github.com/jinzhu/gorm"
	// standart empty sql import for init func
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Database ...
type Database struct {
	Type string
	Path string
	DB   *gorm.DB
}

// LoadDb ...
func LoadDb(d *Database) {
	var err error
	// Database part
	d.DB, err = gorm.Open(d.Type, d.Path)
	if err != nil {
		log.Fatal("Failed open database connection")
	}

	d.DB.LogMode(true)
	// Automigration of tables
	d.DB.AutoMigrate(&m.User{}, &m.Message{})
}
