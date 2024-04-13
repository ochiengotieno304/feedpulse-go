package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ochiengotieno304/feedpulse-go/configs"
	"github.com/ochiengotieno304/feedpulse-go/pkg/models"
)

type FeedHandler struct{}

func (FeedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	country, category, page, per_page := r.URL.Query().Get("country"), r.URL.Query().Get("category"), r.URL.Query().Get("page"), r.URL.Query().Get("per_page")

	if country == "" {
		country = "KE"
	}

	if category == "" {
		category = "SPORTS"
	}

	if page == "" {
		page = "1"
	}

	if per_page == "" {
		per_page = "10"
	}

	perPage, _ := strconv.Atoi(per_page)

	pageInt, _ := strconv.Atoi(page)

	var news *[]models.News

	if err := configs.DB.Where("category LIKE ? AND code = ?", fmt.Sprintf("%s%s%s", "%", category, "%"), country).Limit(perPage).Offset((pageInt - 1) * perPage).Find(&news).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(news); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
}
