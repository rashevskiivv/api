package entity

import (
	"fmt"
)

type User struct {
	ID        *int64   `json:"id"`
	Name      *string  `json:"name"`
	Email     string   `json:"email"`
	Interests []string `json:"interests"`
}

type UserAuth struct {
	ID       *string `json:"id"`
	Name     *string `json:"name"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
}

func (u *UserAuth) Validate() error {
	if u.Email == "" {
		return fmt.Errorf("email cannot be empty")
	}
	return nil
}
