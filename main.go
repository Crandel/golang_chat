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
)

// Config struct
type config struct {
	Database string `json:"Database"`
	Version  string `json:"Version"`
	Host     string `json:"Host"`
	Port     string `json:"Port"`
}

// Config object of config
var Config config

func main() {
	r := router.New()
	configFile, err := ioutil.ReadFile("config.json")
	if err != nil {
		Error.Println("opening config file", err.Error())
	}
	err = json.Unmarshal(configFile, &Config)
	if err != nil {
		Error.Println("parsing json", err.Error())
	}
	Debug.Printf("Start Server version %s on %s:%s logger", Config.Version, Config.Host, Config.Port)
	db, err := gorm.Open("postgres", Config.Database)
	defer db.Close()
	// Automigration of tables
	db.AutoMigrate(&User{}, &Message{})

	r.GET("/", PageMainHandler)
	r.ServeFiles("/static/*filepath", http.Dir("./public"))
	server := fmt.Sprintf("%s:%s", Config.Host, Config.Port)
	log.Fatal(http.ListenAndServe(server, r))
}
