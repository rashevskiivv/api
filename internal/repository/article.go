package repository

import (
	"tax-api/internal/entity"
)

type ArticleRepo struct {
}

type ArticleRepository interface {
	InsertArticle(article entity.Article) error
	ReadArticles(filter entity.Filter) ([]entity.Article, error)
	UpdateArticle(article entity.Article, filter entity.Filter) error
	DeleteArticle(filter entity.Filter) error
}

func (repo *ArticleRepo) InsertArticle(article entity.Article) error {

	return nil
}

func (repo *ArticleRepo) ReadArticles(filter entity.Filter) ([]entity.Article, error) {
	var articles []entity.Article

	return articles, nil
}

func (repo *ArticleRepo) UpdateArticle(article entity.Article, filter entity.Filter) error {

	return nil
}

func (repo *ArticleRepo) DeleteArticle(filter entity.Filter) error {

	return nil
}
