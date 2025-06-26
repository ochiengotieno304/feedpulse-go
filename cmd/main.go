package main

import (
	"log"
	"net/http"

	"github.com/ochiengotieno304/feedpulse-go/api/handlers"
	"github.com/ochiengotieno304/feedpulse-go/internal/middleware"
)

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Healthy"))
}

func main() {
	mux := http.NewServeMux()
	feedHandlers := handlers.NewFeedHandlers()

	readFeedHandler := http.HandlerFunc(feedHandlers.ReadSingleFeedHandler)
	readAllFeads := http.HandlerFunc(feedHandlers.ReadAllFeedsHandler)
	countriesHandler := http.HandlerFunc(handlers.SupportedCountryHandler)

	// AUTHENTICATED
	mux.Handle("GET /api/feeds", middleware.RapidProxySecretCheck(readAllFeads))
	mux.Handle("GET /api/feeds/{id}", middleware.RapidProxySecretCheck(readFeedHandler))
	mux.Handle("GET /api/countries", middleware.RapidProxySecretCheck(countriesHandler))

	// UNAUTHENTICATED
	mux.Handle("GET /health", http.HandlerFunc(healthCheckHandler))

	log.Println("Listening on port 8080")
	http.ListenAndServe(":8080", mux)
}



