package entity

import (
	"fmt"
	"strings"
)

type Answer struct {
	ID         *int64 `json:"id"`
	Answer     string `json:"answer"`
	IDQuestion int64  `json:"id_question"`
	IsRight    bool   `json:"is_right"`
}

func (m *Answer) Validate() error {
	if len(m.Answer) == 0 {
		return fmt.Errorf("answer can not be empty")
	}
	if strings.Contains(strings.ToLower(m.Answer), "drop") {
		return fmt.Errorf("answer contains \"drop\". It is restricted")
	}
	if strings.Contains(strings.ToLower(m.Answer), "delete") {
		return fmt.Errorf("answer contains \"delete\". It is restricted")
	}
	if strings.Contains(m.Answer, "--") {
		return fmt.Errorf("answer contains \"--\". It is restricted")
	}

	if m.IDQuestion == 0 {
		return fmt.Errorf("id_question can not be zero")
	}

	return nil
}
