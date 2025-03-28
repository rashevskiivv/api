package entity

import (
	"fmt"
	"strings"
)

type TestFilter struct {
	ID                 []int64  `json:"id,omitempty"`
	Title              []string `json:"title,omitempty"`
	Description        []string `json:"description,omitempty"`
	AveragePassingTime []string `json:"average_passing_time,omitempty"`
	IDSkill            []int64  `json:"id_skill,omitempty"`
	Limit              uint     `json:"limit,omitempty"`
}

func (f *TestFilter) Validate() error {
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
			if strings.Contains(strings.ToLower(s), "delete") {
				return fmt.Errorf("title contains \"delete\". It is restricted")
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
			if strings.Contains(strings.ToLower(s), "delete") {
				return fmt.Errorf("description contains \"delete\". It is restricted")
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
			if strings.Contains(strings.ToLower(s), "delete") {
				return fmt.Errorf("average_passing_time contains \"delete\". It is restricted")
			}
		}
	}

	if len(f.IDSkill) > 0 {
		for i, v := range f.IDSkill {
			if v == 0 {
				return fmt.Errorf("%v. id_skill can not be zero", i)
			}
		}
	}

	if f.Limit < 0 {
		return fmt.Errorf("limit can not be negative")
	}

	return nil
}

type StartTestInput struct {
	IDTest int64 `json:"id_test"`
	IDUser int64 `json:"id_user"`
}

type StartTestOutput struct {
	NumberOfQuestions int8  `json:"number_of_questions"`
	IDFirstQuestion   int64 `json:"id_first_question"`
}

type EndTestInput struct {
	IDTest int64 `json:"id_test"`
	IDUser int64 `json:"id_user"`
}

type EndTestOutput struct {
	Score             int8 `json:"score"`
	NumberOfQuestions int8 `json:"number_of_questions"`
}
