package question

import (
	"context"
	"fmt"
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

func (r *Repo) Upsert(ctx context.Context, input entity.Question) (*entity.Question, error) {
	var id int64

	const q = `INSERT INTO @table ("question", "id_test")
VALUES (@question, @id_test)
RETURNING id;`
	args := pgx.NamedArgs{
		"table":    entity.TableNameQuestion,
		"question": input.Question,
		"id_test":  input.IDTest,
	}

	err := r.DB.QueryRow(ctx, q, args).Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("unable to insert or update row: %v", err)
	}

	return &entity.Question{ID: &id}, nil
}

func (r *Repo) Read(ctx context.Context, filter entity.QuestionFilter) ([]entity.Question, error) {
	var output []entity.Question

	q := r.builder.Select(
		"id",
		"question",
		"id_test",
	).From(entity.TableNameQuestion)

	// Where
	if len(filter.ID) > 0 {
		q = q.Where(squirrel.Eq{"id": filter.ID})
	}
	if len(filter.Question) > 0 {
		q = q.Where(squirrel.Eq{"question": filter.Question})
	}
	if len(filter.IDTest) > 0 {
		q = q.Where(squirrel.Eq{"id_test": filter.IDTest})
	}

	// Limit
	if filter.Limit != 0 {
		q = q.Limit(uint64(filter.Limit))
	}

	sql, args, err := q.ToSql()
	if err != nil {
		return nil, fmt.Errorf("unable to convert query to sql: %v", err)
	}

	rows, err := r.DB.Query(ctx, sql, args...)
	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("unable to query questions: %v", err)
	}

	for rows.Next() {
		question := entity.Question{}
		err = rows.Scan(
			&question.ID,
			&question.Question,
			&question.IDTest,
		)
		if err != nil {
			return nil, fmt.Errorf("unable to scan row: %v", err)
		}
		output = append(output, question)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return output, nil
}

func (r *Repo) Delete(ctx context.Context, filter entity.QuestionFilter) error {
	q := r.builder.Delete(entity.TableNameQuestion)

	// Where
	if len(filter.ID) > 0 {
		q = q.Where(squirrel.Eq{"id": filter.ID})
	}
	if len(filter.Question) > 0 {
		q = q.Where(squirrel.Eq{"question": filter.Question})
	}
	if len(filter.IDTest) > 0 {
		q = q.Where(squirrel.Eq{"id_test": filter.IDTest})
	}

	// Limit
	if filter.Limit != 0 {
		q = q.Limit(uint64(filter.Limit))
	}

	sql, args, err := q.ToSql()
	if err != nil {
		return fmt.Errorf("unable to convert query to sql: %v", err)
	}

	_, err = r.DB.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("unable to delete questions: %v", err)
	}

	return nil
}
