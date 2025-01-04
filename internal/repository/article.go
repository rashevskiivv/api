package repository

import (
	"tax-api/internal/entity"
)

// todo pgx?

type ArticleRepo struct {
	// todo store db
}

type ArticleRepository interface {
	InsertArticle(article entity.Test) error
	ReadArticles(filter entity.Filter) ([]entity.Test, error)
	UpdateArticle(article entity.Test, filter entity.Filter) error
	DeleteArticle(filter entity.Filter) error
}

func (repo *ArticleRepo) InsertArticle(article entity.Test) error {

	return nil
}

func (repo *ArticleRepo) ReadArticles(filter entity.Filter) ([]entity.Test, error) {
	var articles []entity.Test

	return articles, nil
}

func (repo *ArticleRepo) UpdateArticle(article entity.Test, filter entity.Filter) error {

	return nil
}

func (repo *ArticleRepo) DeleteArticle(filter entity.Filter) error {

	return nil
}
