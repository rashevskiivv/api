package entity

type User struct {
	ID    *int64  `json:"id"`
	Name  *string `json:"name"`
	Email string  `json:"email"`
}

type UserAuth struct {
	ID       *string `json:"id"`
	Name     *string `json:"name"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
}
