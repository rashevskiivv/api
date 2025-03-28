package entity

import "fmt"

type AnswerQuestionFilter struct {
	AF    AnswerFilter
	QF    QuestionFilter
	Limit uint `json:"limit,omitempty"`
}

func (f *AnswerQuestionFilter) Validate() error {
	err := f.AF.Validate()
	if err != nil {
		return err
	}

	err = f.QF.Validate()
	if err != nil {
		return err
	}

	if f.Limit < 0 {
		return fmt.Errorf("limit can not be negative")
	}

	return nil
}

type TestQuestionFilter struct {
	QF    QuestionFilter
	TF    TestFilter
	Limit uint `json:"limit,omitempty"`
}

func (f *TestQuestionFilter) Validate() error {
	err := f.QF.Validate()
	if err != nil {
		return err
	}

	err = f.TF.Validate()
	if err != nil {
		return err
	}

	if f.Limit < 0 {
		return fmt.Errorf("limit can not be negative")
	}

	return nil
}

type TestSkillFilter struct {
	TF    TestFilter
	SF    SkillFilter
	Limit uint `json:"limit,omitempty"`
}

func (f *TestSkillFilter) Validate() error {
	err := f.TF.Validate()
	if err != nil {
		return err
	}

	err = f.SF.Validate()
	if err != nil {
		return err
	}

	if f.Limit < 0 {
		return fmt.Errorf("limit can not be negative")
	}

	return nil
}

type UserSkillFilter struct {
	UF    UserFilter
	SF    SkillFilter
	Limit uint `json:"limit,omitempty"`
}

func (f *UserSkillFilter) Validate() error {
	err := f.UF.Validate()
	if err != nil {
		return err
	}

	err = f.SF.Validate()
	if err != nil {
		return err
	}

	if f.Limit < 0 {
		return fmt.Errorf("limit can not be negative")
	}

	return nil
}

type SkillVacancyFilter struct {
	SF    SkillFilter
	VF    VacancyFilter
	Limit uint `json:"limit,omitempty"`
}

func (f *SkillVacancyFilter) Validate() error {
	err := f.SF.Validate()
	if err != nil {
		return err
	}

	err = f.VF.Validate()
	if err != nil {
		return err
	}

	if f.Limit < 0 {
		return fmt.Errorf("limit can not be negative")
	}

	return nil
}
