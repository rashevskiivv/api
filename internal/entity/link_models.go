package entity

type AnswerQuestion struct {
	A Answer   `json:"answer"`
	Q Question `json:"question"`
}

func (f *AnswerQuestion) Validate() error {
	err := f.A.Validate()
	if err != nil {
		return err
	}
	return f.Q.Validate()
}

type TestQuestion struct {
	Q Question `json:"question"`
	T Test     `json:"test"`
}

func (f *TestQuestion) Validate() error {
	err := f.T.Validate()
	if err != nil {
		return err
	}
	return f.Q.Validate()
}

type TestSkill struct {
	T Test  `json:"test"`
	S Skill `json:"skill"`
}

func (f *TestSkill) Validate() error {
	err := f.T.Validate()
	if err != nil {
		return err
	}
	return f.S.Validate()
}

type UserSkill struct {
	U                User  `json:"user"`
	S                Skill `json:"skill"`
	ProficiencyLevel int   `json:"proficiency_level"`
}

func (f *UserSkill) Validate() error {
	return f.S.Validate()
}

type SkillVacancy struct {
	S Skill   `json:"skill"`
	V Vacancy `json:"vacancy"`
}

func (f *SkillVacancy) Validate() error {
	err := f.V.Validate()
	if err != nil {
		return err
	}
	return f.S.Validate()
}
