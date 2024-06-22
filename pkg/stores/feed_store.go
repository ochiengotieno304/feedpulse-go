package stores

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/ochiengotieno304/feedpulse-go/pkg/db"
	"github.com/ochiengotieno304/feedpulse-go/pkg/models"
)

type FeedStore interface {
	GetAll(r *http.Request) (*[]models.News, error)
	GetSingle(r *http.Request) (*models.News, error)
}

type feedStore struct {
	db  *pgx.Conn
	ctx context.Context
}

func NewFeedStore() FeedStore {
	return &feedStore{
		db:  db.DB(),
		ctx: context.Background(),
	}
}

func (s *feedStore) GetAll(r *http.Request) (*[]models.News, error) {
	var queryString string
	var rows pgx.Rows

	page, err := strconv.ParseInt(r.URL.Query().Get("page"), 10, 32)
	if page == 0 || err != nil {
		page = 1
	}

	pageSize, err := strconv.ParseInt(r.URL.Query().Get("per_page"), 10, 32)
	if pageSize == 0 || err != nil {
		pageSize = 10
	}

	country, category :=
		strings.ToUpper(r.URL.Query().Get("country")),
		strings.ToUpper(r.URL.Query().Get("category"))

	countryPresent := len(country) > 0
	categoryPresent := len(category) > 0

	if countryPresent && categoryPresent {
		queryString = "SELECT id, title, snippet, url, source, code, category, published_date FROM news WHERE code=$1 AND category LIKE '%' || $2 || '%' ORDER BY date DESC LIMIT $3 OFFSET $4;"
		rows, err = s.db.Query(s.ctx, queryString, country, category, pageSize, page)
		if err != nil {
			return nil, err
		}
	} else if countryPresent {
		queryString = "SELECT id, title, snippet, url, source, code, category, published_date FROM news WHERE code=$1 ORDER BY date DESC LIMIT $2 OFFSET $3;"
		rows, err = s.db.Query(s.ctx, queryString, country, pageSize, page)
		if err != nil {
			return nil, err
		}
	} else if categoryPresent {
		queryString = "SELECT id, title, snippet, url, source, code, category, published_date FROM news WHERE category LIKE  '%' || $1 || '%' ORDER BY date DESC LIMIT $2 OFFSET $3;"
		rows, err = s.db.Query(s.ctx, queryString, category, pageSize, page)
		if err != nil {
			return nil, err
		}
	} else {
		country = "KE"
		queryString = "SELECT id, title, snippet, url, source, code, category, published_date FROM news WHERE code=$1 ORDER BY date DESC LIMIT $2 OFFSET $3;"
		rows, err = s.db.Query(s.ctx, queryString, country, pageSize, page)
		if err != nil {
			return nil, err
		}
	}

	defer rows.Close()

	var feeds []models.News
	for rows.Next() {
		var feed models.News
		if err := rows.Scan(
			&feed.ID,
			&feed.Title,
			&feed.Snippet,
			&feed.URL,
			&feed.Source,
			&feed.Code,
			&feed.Category,
			&feed.PublishedDate,
		); err != nil {
			return nil, err
		}

		feeds = append(feeds, feed)
	}

	return &feeds, err
}

func (s *feedStore) GetSingle(r *http.Request) (*models.News, error) {

	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)

	var feed models.News

	if err != nil {
		id = 1
	}

	if err := s.db.QueryRow(s.ctx, "SELECT * FROM news WHERE id=$1", id).Scan(
		&feed.ID,
		&feed.Title,
		&feed.Snippet,
		&feed.URL,
		&feed.Source,
		&feed.Code,
		&feed.Category,
		&feed.PublishedDate,
	); err != nil {
		return nil, err
	}

	return &feed, nil
}
