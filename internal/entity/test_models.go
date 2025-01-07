package entity

// Test model
type Test struct {
	ID                 *int64  `json:"id"`
	Title              string  `json:"title"`
	Description        *string `json:"description"`
	AveragePassingTime *string `json:"average_passing_time"`
	IDSkill            *int64  `json:"id_skill"`
}
