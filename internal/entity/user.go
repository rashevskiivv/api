package entity

// User Модель пользователя
type User struct {
	ID       int64   `json:"id"`
	Name     *string `json:"name"`
	INN      *string `json:"inn"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
}

type Operation struct {
	ID            int64   `json:"id"`
	UserID        int64   `json:"userID"`
	Category      string  `json:"category"`
	AmountOfMoney float64 `json:"amountOfMoney"`
}

// Tax Модель налога
type Tax struct {
	Operation
}

// Income Модель дохода
type Income struct {
	Operation
}

// Expenses Модель траты
type Expenses struct {
	Operation
}
