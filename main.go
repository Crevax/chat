package main

import (
	"log"
	"net/http"
	"os"

	"fmt"

	"github.com/crevax/chat/templating"
)

func main() {
	http.Handle("/", &templating.TemplateHandler{Filepath: []string{"templates", "html", "chat.html"}})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
