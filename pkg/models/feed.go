package models

import "github.com/jackc/pgx/pgtype"

type News struct {
	ID            int64       `json:"id"`
	Title         string      `json:"title"`
	Snippet       string      `json:"snippet"`
	URL           string      `json:"url"`
	Source        string      `json:"source"`
	Code          string      `json:"code"`
	Category      string      `json:"category"`
	PublishedDate pgtype.Date `json:"published_date"`
}

type Response struct {
	Message [16]byte
	Data    [16]byte
	Code    uint16
}
