package main

import (
	"log"
	"net/http"

	"github.com/ochiengotieno304/feedpulse-go/api/handlers"
	"github.com/ochiengotieno304/feedpulse-go/configs"
	"github.com/ochiengotieno304/feedpulse-go/internal/middleware"
)

func init() {
	configs.ConnectDatabase()
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Healthy"))
}

func main() {
	mux := http.NewServeMux()
	feedHandler := handlers.FeedHandler{}
	getSingleFeedHandler := http.HandlerFunc(handlers.GetSingleFeedHandler)

	mux.Handle("GET /api/feeds", middleware.RapidProxySecretCheck(feedHandler))
	mux.Handle("GET /api/feeds/{id}", middleware.RapidProxySecretCheck(getSingleFeedHandler))
	mux.Handle("GET /health", http.HandlerFunc(healthCheckHandler))

	log.Println("Listening on port 8080")
	http.ListenAndServe(":8080", mux)
}
