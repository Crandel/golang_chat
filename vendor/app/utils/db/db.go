package db

import (
	"log"

	"github.com/jinzhu/gorm"
	// standart empty sql import for init func
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB - Database object
var DB *gorm.DB

// Database ...
type Database struct {
	Type string
	Path string
}

// LoadDb ...
func LoadDb(d *Database) {
	var err error
	// Database part
	db, err := gorm.Open(d.Type, d.Path)
	if err != nil {
		log.Fatal("Failed open database connection")
	}
	db.LogMode(true)
	DB = db
}
