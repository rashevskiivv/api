package entity

// Answer model
type Answer struct {
	ID         *int64 `json:"id"`
	Answer     string `json:"answer"`
	IDQuestion int64  `json:"id_question"`
	IsRight    bool   `json:"is_right"`
}
