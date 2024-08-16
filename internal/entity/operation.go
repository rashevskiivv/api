package entity

// Operation cross layers operation entity.
type Operation struct {
	ID            int64  `json:"id"`
	UserID        int64  `json:"user_id"`
	Category      string `json:"category"`
	AmountOfMoney int32  `json:"amount_of_money"`
	IsIncome      bool   `json:"is_income"`
}
