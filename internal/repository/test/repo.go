package test

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

func (r *Repo) Upsert(ctx context.Context, input entity.Test) (*entity.Test, error) {
	log.Println("test upsert started")
	defer log.Println("test upsert done")
	var id int64

	const q = `INSERT INTO @table ("title", "description", "average_passing_time", "id_skill")
VALUES (@title, @description, @average_passing_time, @id_skill)
ON CONFLICT ON CONSTRAINT test_ukey
	DO UPDATE SET title       			= EXCLUDED.title,
				  description			= EXCLUDED.description,
				  average_passing_time	= EXCLUDED.average_passing_time,
				  id_skill 				= EXCLUDED.id_skill
RETURNING id;`
	args := pgx.NamedArgs{
		"table":                entity.TableTest,
		"title":                input.Title,
		"description":          input.Description,
		"average_passing_time": input.AveragePassingTime,
		"id_skill":             input.IDSkill,
	}

	err := r.DB.QueryRow(ctx, q, args).Scan(&id)
	if err != nil {
		log.Printf("unable to insert or update row: %v\n", err)
		return nil, fmt.Errorf("unable to insert or update row: %v", err)
	}

	return &entity.Test{ID: &id}, nil
}

func (r *Repo) Read(ctx context.Context, filter entity.TestFilter) ([]entity.Test, error) {
	log.Println("test read started")
	defer log.Println("test read done")
	var output []entity.Test

	q := r.builder.Select(
		"id",
		"title",
		"description",
		"average_passing_time",
		"id_skill",
	).From(entity.TableTest)

	// Where
	if len(filter.ID) > 0 {
		q = q.Where(squirrel.Eq{"id": filter.ID})
	}
	if len(filter.Title) > 0 {
		q = q.Where(squirrel.Eq{"title": filter.Title})
	}
	if len(filter.Description) > 0 {
		q = q.Where(squirrel.Eq{"description": filter.Description})
	}
	if len(filter.AveragePassingTime) > 0 {
		q = q.Where(squirrel.Eq{"average_passing_time": filter.AveragePassingTime})
	}
	if len(filter.IDSkill) > 0 {
		q = q.Where(squirrel.Eq{"id_skill": filter.IDSkill})
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
		log.Printf("unable to query tests: %v\n", err)
		return nil, fmt.Errorf("unable to query tests: %v", err)
	}

	for rows.Next() {
		test := entity.Test{}
		err = rows.Scan(
			&test.ID,
			&test.Title,
			&test.Description,
			&test.AveragePassingTime,
			&test.IDSkill,
		)
		if err != nil {
			log.Printf("unable to scan row: %v\n", err)
			return nil, fmt.Errorf("unable to scan row: %v", err)
		}
		output = append(output, test)
	}
	if rows.Err() != nil {
		log.Println(rows.Err())
		return nil, rows.Err()
	}

	return output, nil
}

func (r *Repo) Delete(ctx context.Context, filter entity.TestFilter) error {
	log.Println("test delete started")
	defer log.Println("test delete done")
	q := r.builder.Delete(entity.TableTest)

	// Where
	if len(filter.ID) > 0 {
		q = q.Where(squirrel.Eq{"id": filter.ID})
	}
	if len(filter.Title) > 0 {
		q = q.Where(squirrel.Eq{"title": filter.Title})
	}
	if len(filter.Description) > 0 {
		q = q.Where(squirrel.Eq{"description": filter.Description})
	}
	if len(filter.AveragePassingTime) > 0 {
		q = q.Where(squirrel.Eq{"average_passing_time": filter.AveragePassingTime})
	}
	if len(filter.IDSkill) > 0 {
		q = q.Where(squirrel.Eq{"id_skill": filter.IDSkill})
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
		log.Printf("unable to delete tests: %v\n", err)
		return fmt.Errorf("unable to delete tests: %v", err)
	}

	return nil
}

func (r *Repo) Start(ctx context.Context, input entity.StartTestInput) (*entity.StartTestOutput, error) {
	// todo
	panic("implement me")
}

func (r *Repo) End(ctx context.Context, filter entity.EndTestInput) (*entity.EndTestOutput, error) {
	// todo
	panic("implement me")
}
