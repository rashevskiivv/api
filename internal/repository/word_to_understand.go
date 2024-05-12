package repository

import "tax-api/internal/entity"

type WordToUnderstandRepo struct {
}

type WordToUnderstandRepository interface {
	InsertWordToUnderstand(article entity.WordToUnderstand) error
	ReadWordsToUnderstand(filter entity.Filter) ([]entity.WordToUnderstand, error)
	UpdateWordToUnderstand(article entity.WordToUnderstand, filter entity.Filter) error
	DeleteWordToUnderstand(filter entity.Filter) error
}

func (repo *ArticleRepo) InsertWordToUnderstand(word entity.WordToUnderstand) error {

	return nil
}

func (repo *ArticleRepo) ReadWordsToUnderstand(filter entity.Filter) ([]entity.WordToUnderstand, error) {
	var words []entity.WordToUnderstand

	return words, nil
}

func (repo *ArticleRepo) UpdateWordToUnderstand(word entity.WordToUnderstand, filter entity.Filter) error {

	return nil
}

func (repo *ArticleRepo) DeleteWordToUnderstand(filter entity.Filter) error {

	return nil
}
