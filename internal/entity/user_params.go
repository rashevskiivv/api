package entity

import (
	"fmt"
	"strings"
)

type UserFilter struct {
	ID    []int64  `json:"id,omitempty"`
	Name  []string `json:"name,omitempty"`
	Email []string `json:"email,omitempty"`
	Limit uint     `json:"limit,omitempty"`
}

func (f *UserFilter) Validate() error {
	if len(f.ID) > 0 {
		for i, v := range f.ID {
			if v == 0 {
				return fmt.Errorf("%v. id can not be zero", i)
			}
		}
	}

	if len(f.Name) > 0 {
		for i, s := range f.Name {
			if len(s) == 0 {
				return fmt.Errorf("%v. name can not be empty", i)
			}
			if strings.Contains(s, "--") {
				return fmt.Errorf("%v. name contains \"--\". It is restricted", i)
			}
			if strings.Contains(strings.ToLower(s), "drop") {
				return fmt.Errorf("%v. name contains \"drop\". It is restricted", i)
			}
		}
	}

	if len(f.Email) > 0 {
		for i, s := range f.Email {
			if len(s) == 0 {
				return fmt.Errorf("%v. email can not be empty", i)
			}
			if strings.Contains(s, "--") {
				return fmt.Errorf("%v. email contains \"--\". It is restricted", i)
			}
			if strings.Contains(strings.ToLower(s), "drop") {
				return fmt.Errorf("%v. email contains \"drop\". It is restricted", i)
			}
		}
	}

	if f.Limit < 0 {
		return fmt.Errorf("limit can not be negative")
	}

	return nil
}
