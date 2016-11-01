package session

import (
	"net/http"

	"github.com/gorilla/sessions"
)

const name string = "session"

var (
	// Store - store for session
	Store *sessions.CookieStore
)

// Session ...
type Session struct {
	SecretKey     string
	EncryptionKey string
	Options       sessions.Options
}

// InitSession ...
func InitSession(s *Session, domain string) {
	Store = sessions.NewCookieStore([]byte(s.SecretKey))
	Store.Options = &s.Options
}

// CheckUserInSession - check if user in session
func CheckUserInSession(s *sessions.Session) bool {
	return s.Values["id"] == nil
}

// Instance returns a new session, never returns an error
func Instance(r *http.Request) *sessions.Session {
	session, _ := Store.Get(r, name)
	return session
}

// Clear deletes all the current session values
func Clear(s *sessions.Session) {
	for k := range s.Values {
		delete(s.Values, k)
	}
}
