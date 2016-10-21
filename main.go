package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/justinas/alice"
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
	// Store - session store object
	Store *sessions.CookieStore
)

// config struct
type config struct {
	Database database  `json:"Database"`
	Version  string    `json:"Version"`
	Host     string    `json:"Host"`
	Domain   string    `json:"Domain"`
	Port     string    `json:"Port"`
	Template templates `json:"Template"`
	Session  session   `json:"Session"`
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

type session struct {
	SecretKey     string
	EncryptionKey string
}

// CheckErr - function for working with errors
func CheckErr(err error, name string) {
	if err != nil {
		Error.Println(err.Error(), name)
	}
}

// Create new httprouter for ListenAndServe http loop
func routeInit() *mux.Router {
	r := mux.NewRouter()
	// Create base list of middlewares
	baseMidList := []alice.Constructor{LogMiddleware}
	// Create auth list of middlewares, extended from base
	authMidList := []alice.Constructor{}
	copy(authMidList, baseMidList)
	// append from base list
	authMidList = append(authMidList, DisallowAnonMiddleware)
	baseAlice := alice.New(baseMidList...)
	authAlice := alice.New(authMidList...)
	r.Handle("/", authAlice.Then(MainHandler)).Name("home")
	r.Handle("/login", baseAlice.Then(LoginHandler)).Methods("GET", "POST").Name("login")
	r.Handle("/sign", baseAlice.Then(SignHandler)).Methods("GET", "POST").Name("sign")

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./public"))))
	return r
}

func init() {
	// We need to parse config json file into Config struct
	configFile, err := ioutil.ReadFile("config.json")
	err = json.Unmarshal(configFile, &Config)
	CheckErr(err, "Parsing json")
	// Fill template struct from our templates dir
	templ := "%s/%s.%s"
	templateDir := fmt.Sprintf("./%s", Config.Template.Folder)
	// list of files from templates dir
	files, err := ioutil.ReadDir(templateDir)
	CheckErr(err, "Read template Directory")
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
	Store = sessions.NewCookieStore([]byte(Config.Session.SecretKey))
	Store.Options = &sessions.Options{
		Domain:   Config.Domain,
		MaxAge:   3600 * 8, // 8 hours
		HttpOnly: true,
	}
}

func main() {
	// Database part
	Db, _ = gorm.Open(Config.Database.Type, Config.Database.Path)
	defer Db.Close()
	Db.LogMode(true)
	// Automigration of tables
	user := &User{}
	message := &Message{}
	Db.AutoMigrate(user, message)

	r := routeInit()
	server := fmt.Sprintf("%s:%s", Config.Host, Config.Port)
	srv := &http.Server{
		Handler: r,
		Addr:    server,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	Debug.Printf("Start Server version %s on %s", Config.Version, server)
	Error.Println(srv.ListenAndServe())
}
