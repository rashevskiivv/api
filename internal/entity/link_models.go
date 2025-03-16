package entity

type AnswerQuestion struct {
	A Answer   `json:"answer"`
	Q Question `json:"question"`
}

type QuestionTest struct {
	Q Question `json:"question"`
	T Test     `json:"test"`
}

type TestSkill struct {
	T Test  `json:"test"`
	S Skill `json:"skill"`
}

type UserSkill struct {
	U User  `json:"user"`
	S Skill `json:"skill"`
}

type SkillVacancy struct {
	S Skill   `json:"skill"`
	V Vacancy `json:"vacancy"`
}
