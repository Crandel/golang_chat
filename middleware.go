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
		//sess := session.Instance(r)

		// If user is not authenticated, don't allow them to access the page
		/*if sess.Values["id"] == nil {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}*/
		Debug.Println("DisallowAnonMiddleware before")
		next.ServeHTTP(w, r)
		Debug.Println(w)
		Debug.Println("DisallowAnonMiddleware after")
	})
}
