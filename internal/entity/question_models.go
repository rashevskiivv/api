package entity

import (
	"fmt"
	"strings"
)

type Question struct {
	ID       *int64 `json:"id"`
	Question string `json:"question"`
	IDTest   int64  `json:"id_test"`
}

func (m *Question) Validate() error {
	if len(m.Question) == 0 {
		return fmt.Errorf("question can not be empty")
	}
	if strings.Contains(m.Question, "--") {
		return fmt.Errorf("question contains \"--\". It is restricted")
	}
	if strings.Contains(strings.ToLower(m.Question), "drop") {
		return fmt.Errorf("question contains \"drop\". It is restricted")
	}
	if strings.Contains(strings.ToLower(m.Question), "delete") {
		return fmt.Errorf("answer contains \"delete\". It is restricted")
	}

	if m.IDTest == 0 {
		return fmt.Errorf("id_test can not be zero")
	}

	return nil
}
