package entity

// Response Модель ответа сервера
type Response struct {
	Data    any    `json:"data"`
	Message string `json:"message"`
	Errors  string `json:"errors"`
}
