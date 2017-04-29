package main

import (
	"log"
	"net/http"
	"os"

	"fmt"

	"github.com/crevax/chat/templating"
)

func main() {
	r := newRoom()
	http.Handle("/", &templating.TemplateHandler{Filepath: []string{"templates", "html", "chat.html"}})
	http.Handle("/room", r)

	go r.run()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Starting web server on", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
