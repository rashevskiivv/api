package entity

import (
	"fmt"
	"strings"
	"time"
)

type Vacancy struct {
	ID          *int64     `json:"id"`
	Title       string     `json:"title"`
	Grade       *string    `json:"grade"`
	Date        *time.Time `json:"date"`
	Description *string    `json:"description"`
}

func (m *Vacancy) Validate() error {
	if len(m.Title) == 0 {
		return fmt.Errorf("title can not be empty")
	}
	if strings.Contains(strings.ToLower(m.Title), "drop") {
		return fmt.Errorf("title contains \"drop\". It is restricted")
	}
	if strings.Contains(strings.ToLower(m.Title), "delete") {
		return fmt.Errorf("title contains \"delete\". It is restricted")
	}
	if strings.Contains(m.Title, "--") {
		return fmt.Errorf("title contains \"--\". It is restricted")
	}

	if m.Grade != nil {
		if len(*m.Grade) == 0 {
			return fmt.Errorf("grade can not be empty")
		}
		if strings.Contains(strings.ToLower(*m.Grade), "drop") {
			return fmt.Errorf("grade contains \"drop\". It is restricted")
		}
		if strings.Contains(strings.ToLower(*m.Grade), "delete") {
			return fmt.Errorf("grade contains \"delete\". It is restricted")
		}
		if strings.Contains(*m.Grade, "--") {
			return fmt.Errorf("grade contains \"--\". It is restricted")
		}
	}

	if m.Description != nil {
		if len(*m.Description) == 0 {
			return fmt.Errorf("description can not be empty")
		}
		if strings.Contains(strings.ToLower(*m.Description), "drop") {
			return fmt.Errorf("description contains \"drop\". It is restricted")
		}
		if strings.Contains(strings.ToLower(*m.Description), "delete") {
			return fmt.Errorf("description contains \"delete\". It is restricted")
		}
		if strings.Contains(*m.Description, "--") {
			return fmt.Errorf("description contains \"--\". It is restricted")
		}
	}
	return nil
}
