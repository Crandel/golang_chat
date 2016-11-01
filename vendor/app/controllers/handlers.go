package controllers

import (
	m "app/models"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	// 	"github.com/gorilla/sessions"
	// 	valid "github.com/asaskevich/govalidator"
)

var messages []m.Message

// MakeHandler - handler wrapper
func MakeHandler(h func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(h)
}

func getTemlates(name string) (*template.Template, error) {
	return template.ParseFiles("templates/base.gtpl", "templates/main.gtpl") //Config.Template.Root, Config.Template.TemplateMap[name])
}

// Redirect - redirect to named router
func Redirect(r *http.Request, name string) (string, error) {
	url, err := mux.CurrentRoute(r).Subrouter().Get(name).URL()
	return url.String(), err
}

// pageMainHandleFunc handler for main page
func pageMainHandleFunc(w http.ResponseWriter, r *http.Request) {
	// Db.Preload("User").Find(&messages)
	templates, err := getTemlates("main")
	main := template.Must(templates, err)
	main.Execute(w, messages)
}

// MainHandler - for main page
var MainHandler = MakeHandler(pageMainHandleFunc)

// // loginHandleFunc - render template for login page
// func loginHandleFunc(w http.ResponseWriter, r *http.Request) {
// 	sess, err := GetSession(r)
// 	if err != nil {
// 		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
// 		return
// 	}

// 	if r.Method == "GET" {
// 		login := template.Must(getTemlates("login"))
// 		if CheckUserInSession(sess) {
// 			url, err := Redirect(r, "home")
// 			if err != nil {
// 				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
// 				return
// 			}
// 			http.Redirect(w, r, url, http.StatusMovedPermanently)
// 			return
// 		}
// 		login.Execute(w, nil)

// 	} else {
// 		err := r.ParseForm()
// 		// logic part of log in
// 		user := &m.User{
// 			Login:    r.FormValue("login"),
// 			Password: r.FormValue("password"),
// 		}
// 		result, err := valid.ValidateStruct(user)
// 		if err == nil || !result {
// 			err := Db.Where(&m.User{Login: user.Login, Password: user.Password}).First(&user).RecordNotFound()
// 			if !err {
// 				sess.Values["id"] = user.ID
// 				err := sess.Save(r, w)
// 				if err != nil {
// 					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
// 					return
// 				}
// 				url, err := Redirect(r, "home")
// 				if err != nil {
// 					http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
// 					return
// 				}
// 				http.Redirect(w, r, url, http.StatusMovedPermanently)
// 			} else {
// 				url, err := Redirect(r, "login")
// 				if err != nil {
// 					http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
// 					return
// 				}
// 				http.Redirect(w, r, url, http.StatusMovedPermanently)
// 			}
// 		}
// 		return
// 	}
// }

// // LoginHandler ...
// var LoginHandler = MakeHandler(loginHandleFunc)

// // signHandleFunc - render template for sign page
// func signHandleFunc(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "GET" {
// 		templates, err := getTemlates("sign")
// 		sign := template.Must(templates, err)
// 		sign.Execute(w, nil)
// 	} else {
// 		r.ParseForm()
// 		// logic part of sign in
// 		user := &m.User{
// 			Login:    r.FormValue("login"),
// 			Email:    r.FormValue("email"),
// 			Password: r.FormValue("password"),
// 		}
// 		result, err := valid.ValidateStruct(user)
// 		if err == nil || !result {
// 			sess, err := GetSession(r)
// 			if err != nil {
// 				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
// 				return
// 			}
// 			Db.Create(user)
// 			sess.Values["id"] = user.ID
// 			err = sess.Save(r, w)
// 			if err != nil {
// 				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
// 				return
// 			}
// 			url, err := Redirect(r, "home")
// 			if err != nil {
// 				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
// 				return
// 			}
// 			http.Redirect(w, r, url, http.StatusMovedPermanently)
// 		}
// 	}
// 	return
// }

// // SignHandler ...
// var SignHandler = MakeHandler(signHandleFunc)

// // signOutHandleFunc - handle func for signout page
// func signOutHandleFunc(w http.ResponseWriter, r *http.Request) {
// 	sess, err := GetSession(r)
// 	if err != nil {
// 		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
// 		return
// 	}

// 	sess.Options = &sessions.Options{
// 		MaxAge: -1,
// 	}
// 	sess.Values = nil
// 	sess.Save(r, w)
// 	url, err := Redirect(r, "login")
// 	if err != nil {
// 		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
// 		return
// 	}

// 	http.Redirect(w, r, url, http.StatusMovedPermanently)
// 	return
// }

// // SignOutHandler ...
// var SignOutHandler = MakeHandler(signOutHandleFunc)
