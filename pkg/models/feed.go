package models

type News struct {
	ID            int64  `json:"id"`
	Title         string `json:"title"`
	Snippet       string `json:"snippet"`
	URL           string `json:"url"`
	Source        string `json:"source"`
	Code          string `json:"code"`
	Date          string `json:"date"`
	Category      string `json:"category"`
	PublishedDate string `json:"published_date"`
}
