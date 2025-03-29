package skill

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

func (r *Repo) Upsert(ctx context.Context, input entity.Skill) (*entity.Skill, error) {
	log.Println("skill upsert started")
	defer log.Println("skill upsert done")
	var id int64

	const q = `INSERT INTO @table ("title")
VALUES (@title)
ON CONFLICT ON CONSTRAINT skill_ukey
	DO UPDATE SET title	= EXCLUDED.title
RETURNING id;`
	args := pgx.NamedArgs{
		"table": entity.TableSkill,
		"title": input.Title,
	}

	err := r.DB.QueryRow(ctx, q, args).Scan(&id)
	if err != nil {
		log.Printf("unable to insert or update row: %v\n", err)
		return nil, fmt.Errorf("unable to insert or update row: %v", err)
	}

	return &entity.Skill{ID: &id}, nil
}

func (r *Repo) Read(ctx context.Context, filter entity.SkillFilter) ([]entity.Skill, error) {
	log.Println("skill read started")
	defer log.Println("skill read done")
	var output []entity.Skill

	q := r.builder.Select(
		"id",
		"title",
	).From(entity.TableSkill)

	// Where
	if len(filter.ID) > 0 {
		q = q.Where(squirrel.Eq{"id": filter.ID})
	}
	if len(filter.Title) > 0 {
		q = q.Where(squirrel.Eq{"title": filter.Title})
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
		log.Printf("unable to query skills: %v\n", err)
		return nil, fmt.Errorf("unable to query skills: %v", err)
	}

	for rows.Next() {
		skill := entity.Skill{}
		err = rows.Scan(
			&skill.ID,
			&skill.Title,
		)
		if err != nil {
			log.Printf("unable to scan row: %v\n", err)
			return nil, fmt.Errorf("unable to scan row: %v", err)
		}
		output = append(output, skill)
	}
	if rows.Err() != nil {
		log.Println(rows.Err())
		return nil, rows.Err()
	}

	return output, nil
}

func (r *Repo) Delete(ctx context.Context, filter entity.SkillFilter) error {
	log.Println("skill delete started")
	defer log.Println("skill delete done")
	q := r.builder.Delete(entity.TableSkill)

	// Where
	if len(filter.ID) > 0 {
		q = q.Where(squirrel.Eq{"id": filter.ID})
	}
	if len(filter.Title) > 0 {
		q = q.Where(squirrel.Eq{"title": filter.Title})
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
		log.Printf("unable to delete skills: %v\n", err)
		return fmt.Errorf("unable to delete skills: %v", err)
	}

	return nil
}
