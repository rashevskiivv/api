package entity

import (
	"fmt"
	"strings"
	"time"
)

type VacancyFilter struct {
	ID          []int64     `json:"id,omitempty"`
	Title       []string    `json:"title,omitempty"`
	Grade       []string    `json:"grade,omitempty"`
	Date        []time.Time `json:"date,omitempty"`
	Description []string    `json:"description,omitempty"`
	Limit       uint        `json:"limit,omitempty"`
}

func (f *VacancyFilter) Validate() error {
	if len(f.ID) > 0 {
		for i, v := range f.ID {
			if v == 0 {
				return fmt.Errorf("%v. id can not be zero", i)
			}
		}
	}

	if len(f.Title) > 0 {
		for i, s := range f.Title {
			if len(s) == 0 {
				return fmt.Errorf("%v. title can not be empty", i)
			}
			if strings.Contains(strings.ToLower(s), "drop") {
				return fmt.Errorf("%v. title contains \"drop\". It is restricted", i)
			}
			if strings.Contains(strings.ToLower(s), "delete") {
				return fmt.Errorf("%v. title contains \"delete\". It is restricted", i)
			}
			if strings.Contains(s, "--") {
				return fmt.Errorf("%v. title contains \"--\". It is restricted", i)
			}
		}
	}

	if len(f.Grade) > 0 {
		for i, s := range f.Grade {
			if len(s) == 0 {
				return fmt.Errorf("%v. grade can not be empty", i)
			}
			if strings.Contains(strings.ToLower(s), "drop") {
				return fmt.Errorf("%v. grade contains \"drop\". It is restricted", i)
			}
			if strings.Contains(strings.ToLower(s), "delete") {
				return fmt.Errorf("%v. grade contains \"delete\". It is restricted", i)
			}
			if strings.Contains(s, "--") {
				return fmt.Errorf("%v. answer contains \"--\". It is restricted", i)
			}
		}
	}

	if len(f.Description) > 0 {
		for i, s := range f.Description {
			if len(s) == 0 {
				return fmt.Errorf("%v. description can not be empty", i)
			}
			if strings.Contains(strings.ToLower(s), "drop") {
				return fmt.Errorf("%v. description contains \"drop\". It is restricted", i)
			}
			if strings.Contains(strings.ToLower(s), "delete") {
				return fmt.Errorf("%v. description contains \"delete\". It is restricted", i)
			}
			if strings.Contains(s, "--") {
				return fmt.Errorf("%v. description contains \"--\". It is restricted", i)
			}
		}
	}

	if f.Limit < 0 {
		return fmt.Errorf("limit can not be negative")
	}

	return nil
}
