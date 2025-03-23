package entity

import (
	"fmt"
	"strings"
)

type Skill struct {
	ID    *int64 `json:"id"`
	Title string `json:"title"`
}

func (m *Skill) Validate() error {
	if len(m.Title) == 0 {
		return fmt.Errorf("title can not be empty")
	}
	if strings.Contains(m.Title, "--") {
		return fmt.Errorf("title contains \"--\". It is restricted")
	}
	if strings.Contains(strings.ToLower(m.Title), "drop") {
		return fmt.Errorf("title contains \"drop\". It is restricted")
	}
	if strings.Contains(strings.ToLower(m.Title), "delete") {
		return fmt.Errorf("title contains \"delete\". It is restricted")
	}

	return nil
}
