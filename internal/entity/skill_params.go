package entity

import (
	"fmt"
	"strings"
)

type SkillFilter struct {
	ID    []int64  `json:"id,omitempty"`
	Title []string `json:"title,omitempty"`
	Limit int32    `json:"limit,omitempty"`
}

func (f *SkillFilter) Validate() error {
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
			if strings.Contains(s, "--") {
				return fmt.Errorf("%v. title contains \"--\". It is restricted", i)
			}
			if strings.Contains(strings.ToLower(s), "drop") {
				return fmt.Errorf("%v. title contains \"drop\". It is restricted", i)
			}
		}
	}

	if f.Limit < 0 {
		return fmt.Errorf("limit can not be negative")
	}

	return nil
}
