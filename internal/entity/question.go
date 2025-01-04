package entity

type Question struct {
	ID       int64  `json:"id"`
	Question string `json:"question"`
	// todo fk test
}
