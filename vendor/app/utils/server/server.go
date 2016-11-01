package server

import (
	"fmt"
	"log"
	"net/http"
)

// Server ...
type Server struct {
	Host     string
	Domain   string
	Port     int
	UseHTTPS bool
	HTTPS    HTTPS
}

// HTTPS ...
type HTTPS struct {
	Port        int
	Certificate string
	Key         string
}

// Run ...
func Run(h http.Handler, s Server) {
	if s.UseHTTPS {
		startHTTPS(h, s)
	} else {
		startHTTP(h, s)
	}
}

// startHTTPS starts the HTTPS listener
func startHTTPS(h http.Handler, s Server) {
	httpsAddress := fmt.Sprintf("%s:%d", s.Host, s.HTTPS.Port)
	log.Println(fmt.Sprintf("Running HTTPS %s", httpsAddress))

	// Start server with  HTTPS listener
	log.Fatal(http.ListenAndServeTLS(httpsAddress, s.HTTPS.Certificate, s.HTTPS.Key, h))
}

// startHTTP starts the HTTP listener
func startHTTP(h http.Handler, s Server) {
	httpAddress := fmt.Sprintf("%s:%d", s.Host, s.Port)
	log.Println(fmt.Sprintf("Running HTTP %s", httpAddress))

	// Start server with HTTP listener
	log.Fatal(http.ListenAndServe(httpAddress, h))
}
