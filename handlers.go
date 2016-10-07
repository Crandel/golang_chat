package main

import (
	"html/template"
	"net/http"

	router "github.com/julienschmidt/httprouter"
)

var messages []Message

// PageMainHandler handler for main page
func PageMainHandler(w http.ResponseWriter, r *http.Request, _ router.Params) {
	Db.Find(&messages)
	Socket()
	main := template.Must(template.ParseFiles("./templates/base.gotmpl", "./templates/main.gotmpl"))
	main.Execute(w, messages)
}

// LoginHandler ...
func LoginHandler(w http.ResponseWriter, r *http.Request, _ router.Params) {
	login := template.Must(template.ParseFiles("./templates/base.gotmpl", "./templates/login.gotmpl"))
	login.Execute(w, nil)
}
