package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ochiengotieno304/feedpulse-go/pkg/models"
	"github.com/ochiengotieno304/feedpulse-go/pkg/stores"
)

type FeedHandler struct {
}

type newsResponse struct {
	Page     string         `json:"page,omitempty"`
	PageSize string         `json:"page_size,omitempty"`
	Feeds    *[]models.News `json:"feeds,omitempty"`
}

var feedStore = stores.NewFeedStore()

func (FeedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	news, err := feedStore.GetAll(r)

	response := newsResponse{
		Feeds:    news,
		Page:     r.URL.Query().Get("page"),
		PageSize: strconv.Itoa(len(*news)),
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(&response); err != nil {
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
