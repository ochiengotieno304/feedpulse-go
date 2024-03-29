package main

import (
	"fmt"
	"net/http"

	"github.com/ochiengotieno304/feedpulse-go/api/controllers"
	"github.com/ochiengotieno304/feedpulse-go/configs"
)

func init() {
	configs.ConnectDatabase()
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/feeds", controllers.FeedsController{}.GetFeeds)
	fmt.Println("Listening on port 7000")
	http.ListenAndServe(":7000", mux)
}
