package link

import (
	"context"
	"fmt"
	"log"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/rashevskiivv/api/internal/entity"
)

func (r *Repo) UpsertTestUser(ctx context.Context, input entity.TestUser) error {
	log.Println("test-user upsert started")
	defer log.Println("test-user upsert done")

	const q = `INSERT INTO test_user ("id_user", "id_test", "score", "number_of_questions") 
VALUES (@id_user, @id_test, @score, @number_of_questions)
ON CONFLICT ON CONSTRAINT test_user_pkey
	DO UPDATE SET score				  = EXCLUDED.score,
	              number_of_questions = EXCLUDED.number_of_questions;`

	args := pgx.NamedArgs{
		"id_user":             input.U.ID,
		"id_test":             input.T.ID,
		"score":               input.Score,
		"number_of_questions": input.NumberOfQuestions,
	}

	_, err := r.DB.Exec(ctx, q, args)
	if err != nil {
		log.Printf("unable to insert or update row: %v\n", err)
		return fmt.Errorf("unable to insert or update row: %v", err)
	}

	return nil
}

func (r *Repo) ReadTestUser(ctx context.Context, input entity.TestUserFilter) ([]entity.TestUser, error) {
	log.Println("test-user read started")
	defer log.Println("test-user read done")
	var output []entity.TestUser

	q := r.builder.Select("id_user",
		"id_test",
		"score",
		"number_of_questions",
	).From(entity.TableTestUser)

	// Where
	if len(input.UF.ID) > 0 {
		q = q.Where(squirrel.Eq{"id_user": input.UF.ID})
	}
	if len(input.TF.ID) > 0 {
		q = q.Where(squirrel.Eq{"id_test": input.TF.ID})
	}
	if len(input.Score) > 0 {
		q = q.Where(squirrel.Eq{"score": input.Score})
	}
	if len(input.NumberOfQuestions) > 0 {
		q = q.Where(squirrel.Eq{"number_of_questions": input.NumberOfQuestions})
	}

	// Limit
	if input.Limit != 0 {
		q = q.Limit(uint64(input.Limit))
	}

	sql, args, err := q.ToSql()
	if err != nil {
		log.Printf("unable to convert query to sql: %v\n", err)
		return nil, fmt.Errorf("unable to convert query to sql: %v", err)
	}

	rows, err := r.DB.Query(ctx, sql, args...)
	if err != nil {
		log.Printf("unable to query user-skill: %v\n", err)
		return nil, fmt.Errorf("unable to query user-skill: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var link entity.TestUser
		err = rows.Scan(
			&link.U.ID,
			&link.T.ID,
			&link.Score,
			&link.NumberOfQuestions,
		)
		if err != nil {
			log.Printf("unable to scan row: %v\n", err)
			return nil, fmt.Errorf("unable to scan row: %v", err)
		}
		output = append(output, link)
	}
	if rows.Err() != nil {
		log.Println(rows.Err())
		return nil, rows.Err()
	}

	return output, nil
}

func (r *Repo) DeleteTestUser(ctx context.Context, input entity.TestUserFilter) error {
	log.Println("test-user delete started")
	defer log.Println("test-user delete done")

	q := r.builder.Delete(entity.TableTestUser)

	// Where
	if len(input.UF.ID) > 0 {
		q = q.Where(squirrel.Eq{"id_user": input.UF.ID})
	}
	if len(input.TF.ID) > 0 {
		q = q.Where(squirrel.Eq{"id_test": input.TF.ID})
	}
	if len(input.Score) > 0 {
		q = q.Where(squirrel.Eq{"score": input.Score})
	}
	if len(input.NumberOfQuestions) > 0 {
		q = q.Where(squirrel.Eq{"number_of_questions": input.NumberOfQuestions})
	}

	// Limit
	if input.Limit != 0 {
		q = q.Limit(uint64(input.Limit))
	}

	sql, args, err := q.ToSql()
	if err != nil {
		log.Printf("unable to convert query to sql: %v\n", err)
		return fmt.Errorf("unable to convert query to sql: %v", err)
	}

	_, err = r.DB.Exec(ctx, sql, args...)
	if err != nil {
		log.Printf("unable to delete test-user: %v\n", err)
		return fmt.Errorf("unable to delete test-user: %v", err)
	}

	return nil
}
