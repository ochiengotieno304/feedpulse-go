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
	var rows pgx.Rows

	page, err := strconv.ParseInt(r.URL.Query().Get("page"), 10, 32)
	if page == 0 || err != nil {
		page = 1
	}

	pageSize, err := strconv.ParseInt(r.URL.Query().Get("per_page"), 10, 32)
	if err != nil {
		pageSize = 10
	}

	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize < 5:
		pageSize = 5
	}

	country, category :=
		strings.ToUpper(r.URL.Query().Get("country")),
		strings.ToUpper(r.URL.Query().Get("category"))

	countryPresent := len(country) > 0
	categoryPresent := len(category) > 0

	// Initialize the base query
	baseQuery := `
			SELECT id, title, snippet, url, source, code, category, published_date
			FROM news
			WHERE 1=1`

	// Slice to store query parameters
	var queryParams []interface{}
	var queryIndex int

	// Check for country presence
	if countryPresent {
		queryIndex++
		baseQuery += ` AND code = $` + strconv.Itoa(queryIndex)
		queryParams = append(queryParams, country)
	}

	// Check for category presence
	if categoryPresent {
		queryIndex++
		baseQuery += ` AND category LIKE '%' || $` + strconv.Itoa(queryIndex) + ` || '%'`
		queryParams = append(queryParams, category)
	}

	// Add ORDER BY, LIMIT, and OFFSET
	queryIndex++
	baseQuery += ` ORDER BY published_date DESC LIMIT $` + strconv.Itoa(queryIndex)
	queryParams = append(queryParams, pageSize)

	queryIndex++
	baseQuery += ` OFFSET $` + strconv.Itoa(queryIndex)
	queryParams = append(queryParams, page)

	// Execute the query
	rows, err = s.db.Query(s.ctx, baseQuery, queryParams...)
	if err != nil {
		return nil, err
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

	if err := s.db.QueryRow(s.ctx, "SELECT id, title, snippet, url, source, code, category, published_date FROM news WHERE id=$1", id).Scan(
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
