package main

import (
	"log"
	"net/http"

	"github.com/y29-h/Formgram/db"
	"github.com/y29-h/Formgram/handlers"
)

func main() {
	db.Init()

	mux := http.NewServeMux()

	mux.HandleFunc("/register", handlers.RegisterHandler)
	mux.HandleFunc("/login", handlers.LoginHandler)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
