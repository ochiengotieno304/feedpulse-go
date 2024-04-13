package main

import (
	"log"
	"net/http"

	"github.com/ochiengotieno304/feedpulse-go/api/handlers"
	"github.com/ochiengotieno304/feedpulse-go/configs"
)

func init() {
	configs.ConnectDatabase()
}

func main() {
	mux := http.NewServeMux()
	newsController := handlers.FeedHandler{}

	mux.Handle("GET /api/feeds", newsController)

	log.Println("Listening on port 7000")
	http.ListenAndServe(":7000", mux)
	// log.Fatal(http.ListenAndServe(":7000", mux))
}
