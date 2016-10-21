package main

import "net/http"

// MakeHandler - handler wrapper
func MakeHandler(h func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(h)
}

// LogMiddleware - logging handler
func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Debug.Println(r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

// DisallowAnonMiddleware - middleware to disallow anonymous users
func DisallowAnonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get session
		sess := GetSession(r)

		// If user is not authenticated, don't allow them to access the page
		if sess.Values["id"] == nil {
			Redirect(w, r, "login")
		}
		next.ServeHTTP(w, r)
	})
}
