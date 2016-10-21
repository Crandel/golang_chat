package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"

	valid "github.com/asaskevich/govalidator"
)

var (
	messages []Message
)

func getTemlates(name string) (*template.Template, error) {
	return template.ParseFiles(Config.Template.Root, Config.Template.TemplateMap[name])
}

// Redirect - redirect to named router
func Redirect(w http.ResponseWriter, r *http.Request, name string) {
	url, err := mux.CurrentRoute(r).Subrouter().Get(name).URL()
	if err != nil {
		Error.Println(err)
	}
	http.Redirect(w, r, url.String(), 302)
}

// GetSession - return Session pointer
func GetSession(r *http.Request) *sessions.Session {
	sess, err := Store.Get(r, "auth")
	if err != nil {
		Error.Println(err)
		return nil
	}
	return sess
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
		login := template.Must(getTemlates("login"))
		login.Execute(w, nil)
	} else {
		r.ParseForm()
		// logic part of log in
		user := &User{
			Login:    r.FormValue("login"),
			Password: r.FormValue("password"),
		}
		result, err := valid.ValidateStruct(user)
		if err == nil || !result {
			err := Db.Where(&User{Login: user.Login, Password: user.Password}).First(&user).RecordNotFound()
			if !err {
				sess := GetSession(r)
				sess.Values["id"] = user.ID
				sess.Save(r, w)
				Redirect(w, r, "home")
			} else {
				Redirect(w, r, "login")
			}
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
		// logic part of sign in
		user := &User{
			Login:    r.FormValue("login"),
			Email:    r.FormValue("email"),
			Password: r.FormValue("password"),
		}
		result, err := valid.ValidateStruct(user)
		if err == nil || !result {
			sess := GetSession(r)
			Db.Create(user)
			sess.Values["id"] = user.ID
			sess.Save(r, w)
			Redirect(w, r, "home")
		}
	}
}

// SignHandler ...
var SignHandler = MakeHandler(signHandleFunc)
