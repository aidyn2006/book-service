package dto

type BookDTO struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Publisher   string `json:"publisher"`
	ISBN        string `json:"isbn"`
	Genre       string `json:"genre"`
	Stock       int    `json:"stock"`
	TotalPages  int    `json:"total_pages"`
	Language    string `json:"language"`
	PublishedAt string `json:"published_at"`
	Description string `json:"description"`
}
