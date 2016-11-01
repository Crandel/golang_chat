package session

import "github.com/gorilla/sessions"

// Session ...
type Session struct {
	SecretKey     string
	EncryptionKey string
	Store         *sessions.CookieStore
}

// InitSession ...
func InitSession(s *Session, domain string) {
	s.Store = sessions.NewCookieStore([]byte(s.SecretKey))
	s.Store.Options = &sessions.Options{
		Domain:   domain,
		MaxAge:   3600 * 8, // 8 hours
		HttpOnly: true,
	}
}
