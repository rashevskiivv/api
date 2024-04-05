package entity

// Article Модель статьи
type Article struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Text        string `json:"text"`
}
