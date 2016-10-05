package main

import (
	"log"
	"net/http"
	"os"
	"time"

	router "github.com/julienschmidt/httprouter"
)

var (
	// Debug logger
	Debug = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	// Error logger
	Error = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
)

type timeHandler struct {
	format string
}

func check(e error, str string) {
	if e != nil {
		Error.Panic(e.Error(), str)
	}
}

func main() {
	Debug.Print("Start Server on 8080 port logger")
	r := router.New()
	r.GET("/time", func(w http.ResponseWriter, r *http.Request, p router.Params) {
		th := timeHandler{format: time.RFC1123}
		tm := time.Now().Format(th.format)
		Debug.Println(tm)
		Socket()
		_, err := w.Write([]byte("The Datetime is: " + tm))
		check(err, "Write timeHandler")
	})
	err := http.ListenAndServe(":8080", r)
	check(err, "Server")
}
