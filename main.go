package main

import (
	"encoding/json"
	"html/template"
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
type Config struct {
	Database string `json:"Database"`
	Version  string `json:"Version"`
}

// C object of Config
var C Config

func main() {
	Debug.Print("Start Server on 8080 port logger")
	r := router.New()
	configFile, err := ioutil.ReadFile("config.json")
	if err != nil {
		Error.Println("opening config file", err.Error())
	}
	err = json.Unmarshal(configFile, &C)
	if err != nil {
		Error.Println("parsing json", err.Error())
	}
	Debug.Println(C)
	db, err := gorm.Open("postgres", C.Database)
	defer db.Close()
	// Automigration of tables
	db.AutoMigrate(&User{}, &Message{})

	r.GET("/main", func(w http.ResponseWriter, r *http.Request, p router.Params) {
		// firstUser := User{1, "example@mail.com", "first login", "pass1"}
		// secondUser := User{2, "2example@mail.com", "second login", "pass2"}
		// firstMessage := Message{firstUser, "first Message from first user", time.Now()}
		// secondMessage := Message{secondUser, "first Message from second user", time.Now()}
		// messages := []Message{firstMessage, secondMessage}
		// Debug.Println(messages)
		messages := []Message{}
		Socket()
		base := template.Must(template.ParseFiles("./templates/base.gotmpl", "./templates/main.gotmpl"))
		base.Execute(w, messages)
	})
	r.ServeFiles("/static/*filepath", http.Dir("./public"))
	log.Fatal(http.ListenAndServe(":8080", r))
}
