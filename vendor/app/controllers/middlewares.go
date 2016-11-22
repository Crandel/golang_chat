package controllers

import (
	m "app/models"
	s "app/utils/session"
	"context"
	"log"
	"net/http"
)

// LogMiddleware - logging middleware
func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

// UserInContext - save user from session in every request
func UserInContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sess := s.Instance(r)
		if s.CheckUserInSession(sess) {
			id, err := s.GetUserID(sess)
			if err != nil {
				log.Println(err)
				return
			}
			// return user pointer!
			user, err := m.GetUserByID(id)
			if err == nil {
				ctx := context.WithValue(r.Context(), "user", user)
				r = r.WithContext(ctx)
			}
		} else {
			url, err := RedirectFunc("login")
			if err != nil {
				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
				return
			}
			http.Redirect(w, r, url, http.StatusMovedPermanently)
			return
		}
		next.ServeHTTP(w, r)
	})
}
