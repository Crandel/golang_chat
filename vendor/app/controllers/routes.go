package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

// RouteInit - Create new httprouter for ListenAndServe http loop
func RouteInit() *mux.Router {
	r := mux.NewRouter()
	// Create base list of middlewares
	baseMidList := []alice.Constructor{LogMiddleware}
	// Create auth list of middlewares, extended from base
	authMidList := make([]alice.Constructor, len(baseMidList))
	copy(authMidList, baseMidList)

	// append from base list
	authMidList = append(authMidList, DisallowAnonMiddleware)
	// baseAlice := alice.New(baseMidList...)
	authAlice := alice.New(authMidList...)
	r.Handle("/", authAlice.Then(MainHandler)).Name("home")
	// r.Handle("/login", baseAlice.Then(LoginHandler)).Methods("GET", "POST").Name("login")
	// r.Handle("/sign", baseAlice.Then(SignHandler)).Methods("GET", "POST").Name("sign")
	// r.Handle("/signout", baseAlice.Then(SignOutHandler)).Methods("GET").Name("signout")

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./public"))))
	return r
}
