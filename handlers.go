package main

import (
	"html/template"
	"net/http"

	valid "github.com/asaskevich/govalidator"
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
	Debug.Printf("%#v", r.Form)
	// logic part of log in
	user := &User{
		Login:    r.FormValue("login"),
		Password: r.FormValue("password"),
	}
	result, err := valid.ValidateStruct(user)
	if err == nil || !result {
		Db.Create(user)
	}
}
