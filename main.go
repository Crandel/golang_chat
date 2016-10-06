package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	router "github.com/julienschmidt/httprouter"
)

var (
	// Debug logger
	Debug = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	// Error logger
	Error = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
)

func check(e error, str string) {
	if e != nil {
		Error.Panic(e.Error(), str)
	}
}

// Message is struct for template
type Message struct {
	Messages []string
}

func main() {
	Debug.Print("Start Server on 8080 port logger")
	r := router.New()
	r.GET("/main", func(w http.ResponseWriter, r *http.Request, p router.Params) {
		messages := Message{[]string{"first", "sec", "third", "fourth", "Fifth"}}
		Debug.Println(messages)
		Socket()
		base := template.Must(template.ParseFiles("./templates/base.gotmpl", "./templates/main.gotmpl"))
		base.Execute(w, messages)
	})
	r.ServeFiles("/static/*filepath", http.Dir("./public"))
	err := http.ListenAndServe(":8080", r)
	check(err, "Server")
}
