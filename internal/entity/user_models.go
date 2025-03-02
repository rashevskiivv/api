package entity

// User model
type User struct {
	ID    *int64  `json:"id"`
	Name  *string `json:"name"`
	Email string  `json:"email"`
}

// UserSkillLink model
type UserSkillLink struct {
	IDUser           int64 `json:"id_user"`
	IDSkill          int64 `json:"id_skill"`
	ProficiencyLevel int   `json:"proficiency_level"`
}
