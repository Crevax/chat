package main

import (
	"log"
	"net/http"
	"os"

	"fmt"

	"github.com/crevax/chat/templating"
	"github.com/crevax/chat/trace"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/google"
)

func main() {
	r := newRoom()
	r.tracer = trace.New(os.Stdout)
	http.Handle("/chat", MustAuth(&templating.TemplateHandler{Filepath: []string{"templates", "html", "chat.html"}}))
	http.Handle("/login", &templating.TemplateHandler{Filepath: []string{"templates", "html", "login.html"}})
	http.HandleFunc("/auth/", loginHandler)
	http.Handle("/room", r)

	go r.run()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	gomniauth.SetSecurityKey(os.Getenv("OATH_SECRET_KEY"))
	gomniauth.WithProviders(
		google.New(os.Getenv("GOOGLE_OAUTH_ID"), os.Getenv("GOOGLE_OAUTH_SECRET"),
			os.Getenv("APP_ROOT")+"/auth/callback/google"),
	)

	log.Println("Starting web server on", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
