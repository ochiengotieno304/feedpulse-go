package handlers

import (
	"encoding/json"
	"net/http"
)

type FeedHandler struct{}

func (FeedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	feedStore := NewFeedStore()
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
