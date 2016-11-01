package controllers

import (
	"fmt"
	"log"
	"net/http"
)

// LogMiddleware - logging middleware
func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("LogMiddleware")
		log.Println(r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

// DisallowAnonMiddleware - middleware to disallow anonymous users
func DisallowAnonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// // Get session
		// sess, err := h.GetSession(r)
		// if err != nil {
		// 	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		// 	return
		// }
		// // If user is not authenticated, don't allow them to access the page
		// if h.CheckUserInSession(sess) {
		// 	url, err := h.Redirect(r, "login")
		// 	if err != nil {
		// 		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		// 		return
		// 	}
		// 	http.Redirect(w, r, url, http.StatusMovedPermanently)
		// 	return
		// }
		log.Println("DisallowAnonMiddleware")
		next.ServeHTTP(w, r)
	})
}
