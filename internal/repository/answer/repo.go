package answer

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

func (r *Repo) Upsert(ctx context.Context, input entity.Answer) (*entity.Answer, error) {
	log.Println("answer upsert started")
	var id int64

	const q = `INSERT INTO @table ("answer", "id_question", "is_right")
VALUES (@answer, @id_question, @is_right)
RETURNING id;`
	args := pgx.NamedArgs{
		"table":       entity.TableNameAnswer,
		"answer":      input.Answer,
		"id_question": input.IDQuestion,
		"is_right":    input.IsRight,
	}

	err := r.DB.QueryRow(ctx, q, args).Scan(&id)
	if err != nil {
		log.Printf("unable to insert or update row: %v\n", err)
		return nil, fmt.Errorf("unable to insert or update row: %v", err)
	}
	log.Println("answer upsert done")

	return &entity.Answer{ID: &id}, nil
}

func (r *Repo) Read(ctx context.Context, filter entity.AnswerFilter) ([]entity.Answer, error) {
	log.Println("answer read started")
	var output []entity.Answer

	q := r.builder.Select(
		"id",
		"answer",
		"id_question",
		"is_right",
	).From(entity.TableNameAnswer)

	// Where
	if len(filter.ID) > 0 {
		q = q.Where(squirrel.Eq{"id": filter.ID})
	}
	if len(filter.Answer) > 0 {
		q = q.Where(squirrel.Eq{"answer": filter.Answer})
	}
	if len(filter.IDQuestion) > 0 {
		q = q.Where(squirrel.Eq{"id_question": filter.IDQuestion})
	}
	if len(filter.IsRight) > 0 {
		q = q.Where(squirrel.Eq{"is_right": filter.IsRight})
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
		log.Printf("unable to query answers: %v\n", err)
		return nil, fmt.Errorf("unable to query answers: %v", err)
	}

	for rows.Next() {
		answer := entity.Answer{}
		err = rows.Scan(
			&answer.ID,
			&answer.Answer,
			&answer.IDQuestion,
			&answer.IsRight,
		)
		if err != nil {
			log.Printf("unable to scan row: %v\n", err)
			return nil, fmt.Errorf("unable to scan row: %v", err)
		}
		output = append(output, answer)
	}
	if rows.Err() != nil {
		log.Println(rows.Err())
		return nil, rows.Err()
	}
	log.Println("answer read done")

	return output, nil
}

func (r *Repo) Delete(ctx context.Context, filter entity.AnswerFilter) error {
	log.Println("answer delete started")
	q := r.builder.Delete(entity.TableNameAnswer)

	// Where
	if len(filter.ID) > 0 {
		q = q.Where(squirrel.Eq{"id": filter.ID})
	}
	if len(filter.Answer) > 0 {
		q = q.Where(squirrel.Eq{"answer": filter.Answer})
	}
	if len(filter.IDQuestion) > 0 {
		q = q.Where(squirrel.Eq{"id_question": filter.IDQuestion})
	}
	if len(filter.IsRight) > 0 {
		q = q.Where(squirrel.Eq{"is_right": filter.IsRight})
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
		log.Printf("unable to delete answers: %v\n", err)
		return fmt.Errorf("unable to delete answers: %v", err)
	}
	log.Println("answer delete done")

	return nil
}
