package main

import (
	"html/template"
	"net/http"

	valid "github.com/asaskevich/govalidator"
)

var (
	messages []Message
)

func getTemlates(name string) (*template.Template, error) {
	return template.ParseFiles(Config.Template.Root, Config.Template.TemplateMap[name])
}

// pageMainHandleFunc handler for main page
func pageMainHandleFunc(w http.ResponseWriter, r *http.Request) {
	Db.Preload("User").Find(&messages)
	Socket()
	templates, err := getTemlates("main")
	main := template.Must(templates, err)
	main.Execute(w, messages)
}

// MainHandler - for main page
var MainHandler = MakeHandler(pageMainHandleFunc)

// loginHandleFunc - render template for login page
func loginHandleFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		templates, err := getTemlates("login")
		login := template.Must(templates, err)
		login.Execute(w, nil)
	} else {
		r.ParseForm()
		Debug.Printf("%#v", r.Form)
		// logic part of log in
		user := &User{
			Login:    r.FormValue("login"),
			Password: r.FormValue("password"),
		}
		result, err := valid.ValidateStruct(user)
		if err == nil || !result {
			//auth logic
			Debug.Printf("%#v\n%#v", result, user)
		}
	}
}

// LoginHandler ...
var LoginHandler = MakeHandler(loginHandleFunc)

// signHandleFunc - render template for sign page
func signHandleFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		templates, err := getTemlates("sign")
		sign := template.Must(templates, err)
		sign.Execute(w, nil)
	} else {
		r.ParseForm()
		Debug.Printf("%#v", r.Form)
		// logic part of sign in
		user := &User{
			Login:    r.FormValue("login"),
			Email:    r.FormValue("email"),
			Password: r.FormValue("password"),
		}
		result, err := valid.ValidateStruct(user)
		if err == nil || !result {
			Db.Create(user)
		}
	}
}

// SignHandler ...
var SignHandler = MakeHandler(signHandleFunc)
