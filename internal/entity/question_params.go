package entity

import (
	"fmt"
	"strings"
)

type QuestionFilter struct {
	ID       []int64  `json:"id,omitempty"`
	Question []string `json:"question,omitempty"`
	IDTest   []int64  `json:"id_test,omitempty"`
	Limit    int32    `json:"limit,omitempty"`
}

func (f *QuestionFilter) Validate() error {
	if len(f.ID) > 0 {
		for i, v := range f.ID {
			if v == 0 {
				return fmt.Errorf("%v. id can not be zero", i)
			}
		}
	}

	if len(f.Question) > 0 {
		for i, s := range f.Question {
			if len(s) == 0 {
				return fmt.Errorf("%v. question can not be empty", i)
			}
			if strings.Contains(s, "--") {
				return fmt.Errorf("%v. question contains \"--\". It is restricted", i)
			}
			if strings.Contains(strings.ToLower(s), "drop") {
				return fmt.Errorf("%v. question contains \"drop\". It is restricted", i)
			}
		}
	}

	if len(f.IDTest) > 0 {
		for i, v := range f.IDTest {
			if v == 0 {
				return fmt.Errorf("%v. id_test can not be zero", i)
			}
		}
	}

	if f.Limit < 0 {
		return fmt.Errorf("limit can not be negative")
	}

	return nil
}
