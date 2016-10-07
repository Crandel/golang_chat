package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	router "github.com/julienschmidt/httprouter"
)

var (
	// Debug logger
	Debug = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	// Error logger
	Error = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	// Config object of config
	Config config
	// Db - database object
	Db *gorm.DB
)

// Config struct
type config struct {
	Database string `json:"Database"`
	Version  string `json:"Version"`
	Host     string `json:"Host"`
	Port     string `json:"Port"`
}

func init() {
	configFile, err := ioutil.ReadFile("config.json")
	if err != nil {
		Error.Println("opening config file", err.Error())
	}
	err = json.Unmarshal(configFile, &Config)
	if err != nil {
		Error.Println("parsing json", err.Error())
	}
}

func main() {
	Db, _ = gorm.Open("postgres", Config.Database)
	defer Db.Close()
	Db.LogMode(true)
	// Automigration of tables
	user := &User{}
	message := &Message{}
	Db.AutoMigrate(user, message)

	// Create new httprouter for ListenAndServe http loop
	r := router.New()
	Debug.Printf("Start Server version %s on %s:%s", Config.Version, Config.Host, Config.Port)
	user1 := User{
		Email:    "cradlemann@gmail.com",
		Login:    "crandel",
		Password: "pass",
		Messages: []Message{{Message: "First message"}, {Message: "second mess"}},
	}
	Debug.Println(user1)
	Db.Set("gorm:save_associations", true).Create(&user1)
	r.GET("/", PageMainHandler)
	r.ServeFiles("/static/*filepath", http.Dir("./public"))
	server := fmt.Sprintf("%s:%s", Config.Host, Config.Port)
	Error.Println(http.ListenAndServe(server, r))
}
