package entity

import (
	"fmt"
	"strings"
)

type Test struct {
	ID                 *int64  `json:"id"`
	Title              string  `json:"title"`
	Description        *string `json:"description"`
	AveragePassingTime *string `json:"average_passing_time"`
	IDSkill            *int64  `json:"id_skill"`
}

func (m *Test) Validate() error {
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

	if m.Description != nil {
		if len(*m.Description) == 0 {
			return fmt.Errorf("description can not be empty")
		}
		if strings.Contains(*m.Description, "--") {
			return fmt.Errorf("description contains \"--\". It is restricted")
		}
		if strings.Contains(strings.ToLower(*m.Description), "drop") {
			return fmt.Errorf("description contains \"drop\". It is restricted")
		}
		if strings.Contains(strings.ToLower(*m.Description), "delete") {
			return fmt.Errorf("description contains \"delete\". It is restricted")
		}
	}

	if m.AveragePassingTime != nil {
		if len(*m.AveragePassingTime) == 0 {
			return fmt.Errorf("average_passing_time can not be empty")
		}
		if strings.Contains(*m.AveragePassingTime, "--") {
			return fmt.Errorf("average_passing_time contains \"--\". It is restricted")
		}
		if strings.Contains(strings.ToLower(*m.AveragePassingTime), "drop") {
			return fmt.Errorf("average_passing_time contains \"drop\". It is restricted")
		}
		if strings.Contains(strings.ToLower(*m.AveragePassingTime), "delete") {
			return fmt.Errorf("average_passing_time contains \"delete\". It is restricted")
		}
	}

	if m.IDSkill != nil && *m.IDSkill == 0 {
		return fmt.Errorf("id_skill can not be zero")
	}

	return nil
}
