package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ochiengotieno304/feedpulse-go/pkg/stores"
)

type FeedHandler struct {
}

var feedStore = stores.NewFeedStore()

func (FeedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	news, err := feedStore.GetAll(r)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(&news); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
}

func GetSingleFeedHandler(w http.ResponseWriter, r *http.Request) {
	feed, err := feedStore.GetSingle(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(&feed); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
}
