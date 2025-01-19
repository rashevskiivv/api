package entity

import (
	"fmt"
	"strings"
)

type AnswerFilter struct {
	ID         []int64  `json:"id,omitempty"`
	Answer     []string `json:"answer,omitempty"`
	IDQuestion []int64  `json:"id_question,omitempty"`
	IsRight    []bool   `json:"is_right,omitempty"`
	Limit      int32    `json:"limit,omitempty"`
}

func (f *AnswerFilter) Validate() error {
	if len(f.ID) > 0 {
		for i, v := range f.ID {
			if v == 0 {
				return fmt.Errorf("%v. id can not be zero", i)
			}
		}
	}

	if len(f.Answer) > 0 {
		for i, s := range f.Answer {
			if len(s) == 0 {
				return fmt.Errorf("%v. answer can not be empty", i)
			}
			if strings.Contains(strings.ToLower(s), "drop") {
				return fmt.Errorf("%v. answer contains \"drop\". It is restricted", i)
			}
			if strings.Contains(strings.ToLower(s), "delete") {
				return fmt.Errorf("%v. answer contains \"delete\". It is restricted", i)
			}
			if strings.Contains(s, "--") {
				return fmt.Errorf("%v. answer contains \"--\". It is restricted", i)
			}
		}
	}

	if len(f.IDQuestion) > 0 {
		for i, v := range f.ID {
			if v == 0 {
				return fmt.Errorf("%v. id_question can not be zero", i)
			}
		}
	}

	if f.Limit < 0 {
		return fmt.Errorf("limit can not be negative")
	}

	return nil
}
