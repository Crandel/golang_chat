package main

import (
	"html/template"
	"net/http"

	router "github.com/julienschmidt/httprouter"
)

var (
	messages []Message
)

func getTemlates(name string) (*template.Template, error) {
	return template.ParseFiles(Config.Template.Root, Config.Template.TemplateMap[name])
}

// PageMainHandler handler for main page
func PageMainHandler(w http.ResponseWriter, r *http.Request, _ router.Params) {
	Db.Preload("User").Find(&messages)
	Socket()
	templates, err := getTemlates("main")
	main := template.Must(templates, err)
	main.Execute(w, messages)
}

// GetLoginHandler - render template for login page
func GetLoginHandler(w http.ResponseWriter, r *http.Request, _ router.Params) {
	templates, err := getTemlates("login")
	login := template.Must(templates, err)
	login.Execute(w, nil)
}

// PostLoginHandler - parce form, validate and save user
func PostLoginHandler(w http.ResponseWriter, r *http.Request, _ router.Params) {
	r.ParseForm()
	// logic part of log in
	Debug.Printf("%#v", r.Form)
	Debug.Println("login:", r.Form["login"])
	Debug.Println("password:", r.Form["password"])
}
