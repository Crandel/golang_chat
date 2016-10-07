package main

import (
	"html/template"
	"net/http"

	router "github.com/julienschmidt/httprouter"
)

// PageMainHandler handler for main page
func PageMainHandler(w http.ResponseWriter, r *http.Request, _ router.Params) {
	// firstUser := User{1, "example@mail.com", "first login", "pass1"}
	// secondUser := User{2, "2example@mail.com", "second login", "pass2"}
	// firstMessage := Message{firstUser, "first Message from first user", time.Now()}
	// secondMessage := Message{secondUser, "first Message from second user", time.Now()}
	// messages := []Message{firstMessage, secondMessage}
	// Debug.Println(messages)
	messages := []Message{}
	Socket()
	base := template.Must(template.ParseFiles("./templates/base.gotmpl", "./templates/main.gotmpl"))
	base.Execute(w, messages)
}
