package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/ochiengotieno304/feedpulse-go/configs"
	"github.com/ochiengotieno304/feedpulse-go/pkg/models"
)

type FeedStore interface {
	GetAll(r io.ReadCloser) (*[]models.News, error)
}

type feedStore struct{}

func NewFeedStore() FeedStore {
	return &feedStore{}
}

type Body struct {
	Country  string
	Category string
	Page     int32
	PerPage  int32
}

func (s *feedStore) GetAll(r io.ReadCloser) (*[]models.News, error) {
	var body Body

	err := json.NewDecoder(r).Decode(&body)
	if err != nil {
		return nil, err
	}

	country, category, page, perPage := strings.ToUpper(body.Country), strings.ToUpper(body.Category), int(body.Page), int(body.PerPage)

	if page == 0 {
		page = 1
	}

	if perPage == 0 {
		perPage = 10
	}

	if country == "" {
		country = "KE"
	}

	if category == "" {
		category = "NEWS"
	}

	var news *[]models.News

	if err := configs.DB.Order("published_date desc").Where("category LIKE ? AND code = ?", fmt.Sprintf("%s%s%s", "%", category, "%"), country).Limit(perPage).Offset((page - 1) * perPage).Find(&news).Error; err != nil {
		return nil, err
	}

	return news, nil
}
