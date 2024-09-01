package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/ochiengotieno304/feedpulse-go/internal/utils"
	"github.com/ochiengotieno304/feedpulse-go/pkg/models"
	"github.com/ochiengotieno304/feedpulse-go/pkg/stores"
)

type FeedHandler interface {
	ReadAllFeedsHandler(w http.ResponseWriter, r *http.Request)
	ReadSingleFeedHandler(w http.ResponseWriter, r *http.Request)
}

type feedHandler struct{}

func NewFeedHandlers() FeedHandler {
	return &feedHandler{}
}

var feedStore = stores.NewFeedStore()

func (*feedHandler) ReadAllFeedsHandler(w http.ResponseWriter, r *http.Request) {
	page := utils.ValidatePage(r.URL.Query().Get("page"))
	pageSize := utils.ValidatePageSize(r.URL.Query().Get("per_page"))

	country, category, language :=
		strings.ToUpper(r.URL.Query().Get("country")),
		strings.ToUpper(r.URL.Query().Get("category")),
		strings.ToLower(r.URL.Query().Get("language"))

	filters := make(map[string]string, 0)
	filters["country"] = country
	filters["category"] = category
	filters["language"] = language

	feeds, err := feedStore.ReadAll(filters, page, pageSize)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(&feeds); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
}

func (*feedHandler) ReadSingleFeedHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	w.Header().Set("Content-Type", "application/json")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&models.Response{
			Message: utils.ErrorInvalidFeedID.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}

	feeds, err := feedStore.Read(idInt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&models.Response{
			Message: utils.ErrorInvalidFeedID.Error(),
			Code:    http.StatusInternalServerError,
		})
		return
	}

	if err := json.NewEncoder(w).Encode(&feeds); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

}
