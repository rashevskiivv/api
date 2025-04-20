package entity

const (
	TableQuestion = "question"
	TableTest     = "test"
	TableAnswer   = "answer"
	TableUser     = "public.\"user\""
	TableVacancy  = "vacancy"
	TableSkill    = "skill"

	TableSkillVacancy = "skill_vacancy"
	TableUserSkill    = "user_skill"
	TableTestUser     = "test_user"

	PathAnswers   = "/answers"
	PathLinks     = "/links"
	PathQuestions = "/questions"
	PathSkills    = "/skills"
	PathTests     = "/tests"
	PathUsers     = "/users"
	PathVacancies = "/vacancies"

	PathStartTest = "/start"
	PathEndTest   = "/end"

	PathTestSkill    = "/test_skill"
	PathUserSkill    = "/user_skill"
	PathSkillVacancy = "/skill_vacancy"

	PathCheck = "/check"

	// Apps Origin header values

	AppAPI             = "api"
	AppRecommendations = "recommendations"
	AppAuth            = "auth"
)
