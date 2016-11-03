package controllers

import (
	"app/models"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

// Templates ...
type Templates struct {
	Root        string
	Ext         string
	Folder      string
	TemplateMap map[string]string
}

var templ *Templates

// RouteInit - Create new httprouter for ListenAndServe http loop
func RouteInit() *mux.Router {
	models.Automigrate()
	r := mux.NewRouter()
	// Create base list of middlewares
	baseMidList := []alice.Constructor{LogMiddleware}
	// Create auth list of middlewares, extended from base
	authMidList := make([]alice.Constructor, len(baseMidList))
	copy(authMidList, baseMidList)

	// append from base list
	authMidList = append(authMidList, DisallowAnonMiddleware)
	baseAlice := alice.New(baseMidList...)
	authAlice := alice.New(authMidList...)
	r.Handle("/", authAlice.Then(MainHandler)).Name("home")
	r.Handle("/login", baseAlice.Then(LoginHandler)).Methods("GET", "POST").Name("login")
	r.Handle("/sign", baseAlice.Then(SignHandler)).Methods("GET", "POST").Name("sign")
	r.Handle("/signout", baseAlice.Then(SignOutHandler)).Methods("GET").Name("signout")

	r.PathPrefix("/static/").Handler(baseAlice.Then(http.StripPrefix("/static/", http.FileServer(http.Dir("./public")))))
	return r
}

// LoadTemplates ...
func LoadTemplates(t *Templates) {
	// Fill template struct from our templates dir
	template := "%s/%s.%s"
	templateDir := fmt.Sprintf("./%s", t.Folder)
	// list of files from templates dir
	files, err := ioutil.ReadDir(templateDir)
	if err != nil {
		log.Fatalln("Read template Directory", err)
	}
	// Set template Root file, all templates will Parse with it
	t.Root = fmt.Sprintf(template, t.Folder, t.Root, t.Ext)
	// Create map with template names
	templates := make(map[string]string)
	for _, f := range files {
		ext := fmt.Sprintf(".%s", t.Ext)
		fullName := f.Name()
		index := strings.Index(fullName, ext)
		// name without extension
		name := fullName[:index]
		templates[name] = fmt.Sprintf("./%s/%s", t.Folder, fullName)
	}
	t.TemplateMap = templates
	templ = t
}
