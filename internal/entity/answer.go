package entity

type Answer struct {
	ID      int64  `json:"id"`
	Answer  string `json:"answer"`
	IsRight bool   `json:"is_right"`
	// todo fk question
}
