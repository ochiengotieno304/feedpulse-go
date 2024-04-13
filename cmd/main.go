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
	feedHandler := handlers.FeedHandler{}

	mux.Handle("GET /api/feeds", feedHandler)

	log.Println("Listening on port 7000")
	http.ListenAndServe(":7000", mux)
}
