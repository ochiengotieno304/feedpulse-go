package models

type News struct {
	ID            int64  `json:"id,omitempty"`
	Title         string `json:"title,omitempty"`
	Snippet       string `json:"snippet,omitempty"`
	URL           string `json:"url,omitempty"`
	Source        string `json:"source,omitempty"`
	Code          string `json:"country,omitempty"`
	Category      string `json:"category,omitempty"`
	Language      string `json:"language,omitempty"`
	PublishedDate string `json:"published_date,omitempty"`
}

type Response struct {
	Message string
	Code    uint16
}
