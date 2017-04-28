package main

import (
	"log"
	"net/http"

	"github.com/crevax/chat/templating"
)

func main() {
	http.Handle("/", &templating.TemplateHandler{Filepath: []string{"templates", "html", "chat.html"}})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
