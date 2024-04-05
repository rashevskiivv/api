package entity

// WordToUnderstand Модель термина
type WordToUnderstand struct {
	ID         int64              `json:"id"`
	Word       string             `json:"word"`
	Definition string             `json:"definition"`
	Synonyms   []WordToUnderstand `json:"synonyms"`
	Antonyms   []WordToUnderstand `json:"antonyms"`
}
