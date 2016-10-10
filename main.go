package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

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

// config struct
type config struct {
	Database database  `json:"Database"`
	Version  string    `json:"Version"`
	Host     string    `json:"Host"`
	Port     string    `json:"Port"`
	Template templates `json:"Template"`
}

// Struct for templates
type templates struct {
	Root        string
	Ext         string
	Folder      string
	TemplateMap map[string]string
}

type database struct {
	Type string
	Path string
}

func init() {
	// We need to parse config json file into Config struct
	configFile, err := ioutil.ReadFile("config.json")
	if err != nil {
		Error.Println("Opening config file", err.Error())
	}
	err = json.Unmarshal(configFile, &Config)
	if err != nil {
		Error.Println("Parsing json", err.Error())
	}
	// Fill template struct from our templates dir
	templ := "%s/%s.%s"
	templateDir := fmt.Sprintf("./%s", Config.Template.Folder)
	// list of files from templates dir
	files, err := ioutil.ReadDir(templateDir)
	if err != nil {
		Error.Println("Read template Directory", err.Error())
	}
	// Set template Root file, all templates will Parse with it
	Config.Template.Root = fmt.Sprintf(templ, Config.Template.Folder, Config.Template.Root, Config.Template.Ext)
	// Create map with template names
	templates := make(map[string]string)
	for _, f := range files {
		ext := fmt.Sprintf(".%s", Config.Template.Ext)
		fullName := f.Name()
		index := strings.Index(fullName, ext)
		// name without extension
		name := fullName[:index]
		templates[name] = fmt.Sprintf("./%s/%s", Config.Template.Folder, fullName)
	}
	Config.Template.TemplateMap = templates
}

func main() {
	Db, _ = gorm.Open(Config.Database.Type, Config.Database.Path)
	defer Db.Close()
	Db.LogMode(true)
	// Automigration of tables
	user := &User{}
	message := &Message{}
	Db.AutoMigrate(user, message)

	// Create new httprouter for ListenAndServe http loop
	r := router.New()
	Debug.Printf("Start Server version %s on %s:%s", Config.Version, Config.Host, Config.Port)
	// user1 := User{
	// 	Email:    "cradlemann@gmail.com",
	// 	Login:    "crandel",
	// 	Password: "pass",
	// 	Messages: []Message{{Message: "First message"}, {Message: "second mess"}},
	// }
	// Debug.Println(user1)
	// Db.Set("gorm:save_associations", true).Create(&user1)
	r.GET("/", PageMainHandler)
	r.GET("/login", GetLoginHandler)
	r.POST("/login", PostLoginHandler)
	r.ServeFiles("/static/*filepath", http.Dir("./public"))
	server := fmt.Sprintf("%s:%s", Config.Host, Config.Port)
	Error.Println(http.ListenAndServe(server, r))
}
