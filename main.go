package main

import (
	"log"
	"net/http"

	"github.com/y29-h/Formgram/db"
)

func main() {
	db.Init()

	mux := http.NewServeMux()

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
