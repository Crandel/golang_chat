package controllers

import (
	m "app/models"
	s "app/utils/session"
	"html/template"
	"log"
	"net/http"

	valid "github.com/asaskevich/govalidator"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// MakeHandler - handler wrapper
func MakeHandler(h func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(h)
}

func getTemlates(name string) (*template.Template, error) {
	return template.ParseFiles(templ.Root, templ.TemplateMap[name])
}

// pageMainHandleFunc handler for main page
func pageMainHandleFunc(w http.ResponseWriter, r *http.Request) {
	messages := []m.Message{}
	// m.GetMessages(&messages)
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
		err := r.ParseForm()
		// logic part of log in
		user := &m.User{
			Login:    r.FormValue("login"),
			Password: r.FormValue("password"),
		}
		result, err := valid.ValidateStruct(user)
		if err == nil || !result {
			err := user.GetUserByLoginPass(user.Login, user.Password)
			if !err {
				sess := s.Instance(r)
				sess.Values["id"] = user.ID
				err := sess.Save(r, w)
				if err != nil {
					log.Println(err)
					return
				}
				url, err := RedirectFunc("home")
				if err != nil {
					http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
					return
				}
				http.Redirect(w, r, url, http.StatusMovedPermanently)
			} else {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			}
			return
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
		user := &m.User{
			Login:    r.FormValue("login"),
			Email:    r.FormValue("email"),
			Password: r.FormValue("password"),
		}
		result, err := valid.ValidateStruct(user)
		if err == nil || !result {
			sess := s.Instance(r)
			if err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
			user.CreateUser()
			sess.Values["id"] = user.ID
			err = sess.Save(r, w)
			if err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
			url, err := RedirectFunc("home")
			if err != nil {
				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
				return
			}
			http.Redirect(w, r, url, http.StatusMovedPermanently)
		}
	}
	return
}

// SignHandler ...
var SignHandler = MakeHandler(signHandleFunc)

// signOutHandleFunc - handle func for signout page
func signOutHandleFunc(w http.ResponseWriter, r *http.Request) {
	sess := s.Instance(r)
	s.Clear(sess)
	url, err := RedirectFunc("login")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	http.Redirect(w, r, url, http.StatusMovedPermanently)
}

// SignOutHandler ...
var SignOutHandler = MakeHandler(signOutHandleFunc)

// NotFoundHandleFunc ...
func NotFoundHandleFunc(w http.ResponseWriter, r *http.Request) {
	log.Println("Not found handler")
	templates, err := getTemlates("404")
	notFound := template.Must(templates, err)
	notFound.Execute(w, nil)
}

// wsHandleFunc ...
func wsHandleFunc(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	sess := s.Instance(r)
	userID, err := s.GetUserID(sess)
	if err != nil {
		log.Println(err)
		return
	}
	user := &m.User{}
	user.GetUserByID(userID)
	hub := GetHub()
	client := &Client{*user, conn, make(chan []byte, 256)}
	hub.register <- client
	go client.write()
	client.read()
}

var WsHandler = MakeHandler(wsHandleFunc)
