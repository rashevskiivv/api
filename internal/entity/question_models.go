package entity

// Question model
type Question struct {
	ID       *int64 `json:"id"`
	Question string `json:"question"`
	IDTest   int64  `json:"id_test"`
}
