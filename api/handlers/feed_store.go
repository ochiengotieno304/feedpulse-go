package handlers

import (
	// "encoding/json"
	"fmt"
	"strconv"
	// "io"
	"net/http"
	"strings"

	"github.com/ochiengotieno304/feedpulse-go/configs"
	"github.com/ochiengotieno304/feedpulse-go/pkg/models"
)

type FeedStore interface {
	GetAll(r *http.Request) (*[]models.News, error)
}

type feedStore struct{}

func NewFeedStore() FeedStore {
	return &feedStore{}
}

func (s *feedStore) GetAll(r *http.Request) (*[]models.News, error) {
	page, err := strconv.ParseInt(r.URL.Query().Get("page"), 10, 32)
	if page == 0 || err != nil {
		page = 1
	}

	pageSize, err := strconv.ParseInt(r.URL.Query().Get("per_page"), 10, 32)
	if pageSize == 0 || err != nil {
		pageSize = 10
	}

	country, category, page, perPage :=
		strings.ToUpper(r.URL.Query().Get("country")),
		strings.ToUpper(r.URL.Query().Get("category")),
		page,
		pageSize

	if country == "" {
		country = "KE"
	}

	if category == "" {
		category = "NEWS"
	}

	var news *[]models.News

	if err := configs.DB.Order("published_date desc").Where("category LIKE ? AND code = ?", fmt.Sprintf("%s%s%s", "%", category, "%"), country).Limit(int(perPage)).Offset(int((page - 1) * perPage)).Find(&news).Error; err != nil {
		return nil, err
	}

	return news, nil
}
