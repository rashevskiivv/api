package question

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

func (r *Repo) Upsert(ctx context.Context, input entity.Question) (*entity.Question, error) {
	log.Println("question upsert started")
	defer log.Println("question upsert done")
	var id int64

	const q = `INSERT INTO @table ("question", "id_test")
VALUES (@question, @id_test)
ON CONFLICT ON CONSTRAINT question_ukey
	DO UPDATE SET question	= EXCLUDED.question,
				  id_test	= EXCLUDED.id_test
RETURNING id;`
	args := pgx.NamedArgs{
		"table":    entity.TableQuestion,
		"question": input.Question,
		"id_test":  input.IDTest,
	}

	err := r.DB.QueryRow(ctx, q, args).Scan(&id)
	if err != nil {
		log.Printf("unable to insert or update row: %v\n", err)
		return nil, fmt.Errorf("unable to insert or update row: %v", err)
	}

	return &entity.Question{ID: &id}, nil
}

func (r *Repo) Read(ctx context.Context, filter entity.QuestionFilter) ([]entity.Question, error) {
	log.Println("question read started")
	defer log.Println("question read done")
	var output []entity.Question

	q := r.builder.Select(
		"id",
		"question",
		"id_test",
	).From(entity.TableQuestion)

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
		log.Printf("unable to convert query to sql: %v\n", err)
		return nil, fmt.Errorf("unable to convert query to sql: %v", err)
	}

	rows, err := r.DB.Query(ctx, sql, args...)
	defer rows.Close()
	if err != nil {
		log.Printf("unable to query questions: %v\n", err)
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
			log.Printf("unable to scan row: %v\n", err)
			return nil, fmt.Errorf("unable to scan row: %v", err)
		}
		output = append(output, question)
	}
	if rows.Err() != nil {
		log.Println(rows.Err())
		return nil, rows.Err()
	}

	return output, nil
}

func (r *Repo) Delete(ctx context.Context, filter entity.QuestionFilter) error {
	log.Println("question delete started")
	defer log.Println("question delete done")
	q := r.builder.Delete(entity.TableQuestion)

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
		log.Printf("unable to convert query to sql: %v\n", err)
		return fmt.Errorf("unable to convert query to sql: %v", err)
	}

	_, err = r.DB.Exec(ctx, sql, args...)
	if err != nil {
		log.Printf("unable to delete questions: %v\n", err)
		return fmt.Errorf("unable to delete questions: %v", err)
	}

	return nil
}
