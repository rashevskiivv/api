package entity

import (
	"fmt"
	"strconv"
	"strings"
)

type TestFilter struct {
	ID                 []string `json:"id,omitempty"`
	Title              []string `json:"title,omitempty"`
	Description        []string `json:"description,omitempty"`
	AveragePassingTime []string `json:"average_passing_time,omitempty"`
	IDSkill            []string `json:"id_skill,omitempty"`
	Limit              int32    `json:"limit,omitempty"`
}

func (f *TestFilter) Validate() error {
	if len(f.ID) > 0 {
		for i, s := range f.ID {
			if len(s) == 0 {
				return fmt.Errorf("%v. id can not be empty", i)
			}
			_, err := strconv.Atoi(s)
			if err != nil {
				return fmt.Errorf("%v. id is not integer", i)
			}
		}
	}

	if len(f.Title) > 0 {
		for i, s := range f.Title {
			if len(s) == 0 {
				return fmt.Errorf("%v. title can not be empty", i)
			}
			if strings.Contains(s, "--") {
				return fmt.Errorf("%v. title contains \"--\". It is restricted", i)
			}
			if strings.Contains(strings.ToLower(s), "drop") {
				return fmt.Errorf("%v. title contains \"drop\". It is restricted", i)
			}
		}
	}

	if len(f.Description) > 0 {
		for i, s := range f.Description {
			if len(s) == 0 {
				return fmt.Errorf("%v. description can not be empty", i)
			}
			if strings.Contains(s, "--") {
				return fmt.Errorf("%v. description contains \"--\". It is restricted", i)
			}
			if strings.Contains(strings.ToLower(s), "drop") {
				return fmt.Errorf("%v. description contains \"drop\". It is restricted", i)
			}
		}
	}

	if len(f.AveragePassingTime) > 0 {
		for i, s := range f.AveragePassingTime {
			if len(s) == 0 {
				return fmt.Errorf("%v. average_passing_time can not be empty", i)
			}
			if strings.Contains(s, "--") {
				return fmt.Errorf("%v. average_passing_time contains \"--\". It is restricted", i)
			}
			if strings.Contains(strings.ToLower(s), "drop") {
				return fmt.Errorf("%v. average_passing_time contains \"drop\". It is restricted", i)
			}
		}
	}

	if len(f.IDSkill) > 0 {
		for i, s := range f.IDSkill {
			if len(s) == 0 {
				return fmt.Errorf("%v. id_skill can not be empty", i)
			}
			_, err := strconv.Atoi(s)
			if err != nil {
				return fmt.Errorf("%v. id_skill is not integer", i)
			}
		}
	}

	if f.Limit < 0 {
		return fmt.Errorf("limit can not be negative")
	}

	return nil
}
