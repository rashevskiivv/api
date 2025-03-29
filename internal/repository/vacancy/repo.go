package vacancy

import (
	"context"
	"fmt"
	"log"
	"tax-api/internal/entity"
	"tax-api/internal/repository"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

type Repo struct {
	*repository.Postgres
	builder squirrel.StatementBuilderType
}

func NewRepo(pg *repository.Postgres) *Repo {
	return &Repo{
		Postgres: pg,
		builder:  squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}

func (r *Repo) Upsert(ctx context.Context, input entity.Vacancy) (*entity.Vacancy, error) {
	log.Println("vacancy upsert started")
	defer log.Println("vacancy upsert done")
	var id int64

	const q = `INSERT INTO @table ("title", "grade", "date", "description")
VALUES (@title, @grade, @date, @description)
ON CONFLICT ON CONSTRAINT vacancy_ukey
	DO UPDATE SET title       	= EXCLUDED.title,
				  grade			= EXCLUDED.grade,
				  "date"		= EXCLUDED.date,
				  description	= EXCLUDED.description
RETURNING id;`
	args := pgx.NamedArgs{
		"table":       entity.TableVacancy,
		"title":       input.Title,
		"grade":       input.Grade,
		"date":        input.Date,
		"description": input.Description,
	}

	err := r.DB.QueryRow(ctx, q, args).Scan(&id)
	if err != nil {
		log.Printf("unable to insert or update row: %v\n", err)
		return nil, fmt.Errorf("unable to insert or update row: %v", err)
	}

	return &entity.Vacancy{ID: &id}, nil
}

func (r *Repo) Read(ctx context.Context, filter entity.VacancyFilter) ([]entity.Vacancy, error) {
	log.Println("vacancy read started")
	defer log.Println("vacancy read done")
	var output []entity.Vacancy

	q := r.builder.Select(
		"id",
		"title",
		"grade",
		"date",
		"description",
	).From(entity.TableVacancy)

	// Where
	if len(filter.ID) > 0 {
		q = q.Where(squirrel.Eq{"id": filter.ID})
	}
	if len(filter.Title) > 0 {
		q = q.Where(squirrel.Eq{"title": filter.Title})
	}
	if len(filter.Grade) > 0 {
		q = q.Where(squirrel.Eq{"grade": filter.Grade})
	}
	if len(filter.Date) > 0 {
		q = q.Where(squirrel.Eq{"date": filter.Date})
	}
	if len(filter.Description) > 0 {
		q = q.Where(squirrel.Eq{"description": filter.Description})
	}

	// Limit
	if filter.Limit != 0 {
		q = q.Limit(uint64(filter.Limit))
	}

	sql, args, err := q.ToSql()
	if err != nil {
		log.Printf("unable to convert query to sql: %v\n", err)
		return nil, fmt.Errorf("unable to convert query to sql: %v", err)
	}

	rows, err := r.DB.Query(ctx, sql, args...)
	defer rows.Close()
	if err != nil {
		log.Printf("unable to query vacancies: %v\n", err)
		return nil, fmt.Errorf("unable to query vacancies: %v", err)
	}

	for rows.Next() {
		vacancy := entity.Vacancy{}
		err = rows.Scan(
			&vacancy.ID,
			&vacancy.Title,
			&vacancy.Grade,
			&vacancy.Date,
			&vacancy.Description,
		)
		if err != nil {
			log.Printf("unable to scan row: %v\n", err)
			return nil, fmt.Errorf("unable to scan row: %v", err)
		}
		output = append(output, vacancy)
	}
	if rows.Err() != nil {
		log.Println(rows.Err())
		return nil, rows.Err()
	}

	return output, nil
}

func (r *Repo) Delete(ctx context.Context, filter entity.VacancyFilter) error {
	log.Println("vacancy delete started")
	defer log.Println("vacancy delete done")
	q := r.builder.Delete(entity.TableVacancy)

	// Where
	if len(filter.ID) > 0 {
		q = q.Where(squirrel.Eq{"id": filter.ID})
	}
	if len(filter.Title) > 0 {
		q = q.Where(squirrel.Eq{"title": filter.Title})
	}
	if len(filter.Grade) > 0 {
		q = q.Where(squirrel.Eq{"grade": filter.Grade})
	}
	if len(filter.Date) > 0 {
		q = q.Where(squirrel.Eq{"date": filter.Date})
	}
	if len(filter.Description) > 0 {
		q = q.Where(squirrel.Eq{"description": filter.Description})
	}

	// Limit
	if filter.Limit != 0 {
		q = q.Limit(uint64(filter.Limit))
	}

	sql, args, err := q.ToSql()
	if err != nil {
		log.Printf("unable to convert query to sql: %v\n", err)
		return fmt.Errorf("unable to convert query to sql: %v", err)
	}

	_, err = r.DB.Exec(ctx, sql, args...)
	if err != nil {
		log.Printf("unable to delete vacancies: %v\n", err)
		return fmt.Errorf("unable to delete vacancies: %v", err)
	}

	return nil
}
