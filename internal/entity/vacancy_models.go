package entity

import "time"

// Vacancy model
type Vacancy struct {
	ID          *int64     `json:"id"`
	Title       string     `json:"title"`
	Grade       *string    `json:"grade"`
	Date        *time.Time `json:"date"`
	Description *string    `json:"description"`
}
