package main

import (
	"html/template"
	"net/http"

	router "github.com/julienschmidt/httprouter"
)

var (
	messages []Message
)

// PageMainHandler handler for main page
func PageMainHandler(w http.ResponseWriter, r *http.Request, _ router.Params) {
	Db.Preload("User").Find(&messages)
	Socket()
	main := template.Must(template.ParseFiles(Config.Template.Root, Config.Template.TemplateMap["main"]))
	main.Execute(w, messages)
}

// GetLoginHandler - render template for login page
func GetLoginHandler(w http.ResponseWriter, r *http.Request, _ router.Params) {
	login := template.Must(template.ParseFiles(Config.Template.Root, Config.Template.TemplateMap["login"]))
	login.Execute(w, nil)
}

// PostLoginHandler - parce form, validate and save user
func PostLoginHandler(w http.ResponseWriter, r *http.Request, p router.Params) {
	// validate form here
}
