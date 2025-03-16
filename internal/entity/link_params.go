package entity

type AnswerQuestionFilter struct {
	AF    AnswerFilter
	QF    QuestionFilter
	Limit uint `json:"limit,omitempty"`
}

type QuestionTestFilter struct {
	QF    QuestionFilter
	TF    TestFilter
	Limit uint `json:"limit,omitempty"`
}

type TestSkillFilter struct {
	TF    TestFilter
	SF    SkillFilter
	Limit uint `json:"limit,omitempty"`
}

type UserSkillFilter struct {
	UF    UserFilter
	SF    SkillFilter
	Limit uint `json:"limit,omitempty"`
}

type SkillVacancyFilter struct {
	SF    SkillFilter
	VF    VacancyFilter
	Limit uint `json:"limit,omitempty"`
}
