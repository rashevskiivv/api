package link

import (
	"context"
	"fmt"
	"log"

	"github.com/rashevskiivv/api/internal/entity"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

func (r *Repo) UpsertSkillVacancy(ctx context.Context, input entity.SkillVacancy) error {
	log.Println("skill-vacancy upsert started")
	defer log.Println("skill-vacancy upsert done")

	const q = `INSERT INTO skill_vacancy ("id_vacancy", "id_skill") VALUES (@id_vacancy, @id_skill);`

	args := pgx.NamedArgs{
		"id_vacancy": input.V.ID,
		"id_skill":   input.S.ID,
	}

	_, err := r.DB.Exec(ctx, q, args)
	if err != nil {
		log.Printf("unable to insert or update row: %v\n", err)
		return fmt.Errorf("unable to insert or update row: %v", err)
	}

	return nil
}

func (r *Repo) ReadSkillVacancy(ctx context.Context, input entity.SkillVacancyFilter) ([]entity.SkillVacancy, error) {
	log.Println("skill-vacancy read started")
	defer log.Println("skill-vacancy read done")
	var output []entity.SkillVacancy

	q := r.builder.Select(
		"id_vacancy",
		"id_skill",
	).From(entity.TableSkillVacancy)

	// Where
	// Skill part
	if len(input.SF.ID) > 0 {
		q = q.Where(squirrel.Eq{"id_skill": input.SF.ID})
	}
	// Vacancy part
	if len(input.VF.ID) > 0 {
		q = q.Where(squirrel.Eq{"id_vacancy": input.VF.ID})
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
	defer rows.Close()
	if err != nil {
		log.Printf("unable to query skill-vacancy: %v\n", err)
		return nil, fmt.Errorf("unable to query skill-vacancy: %v", err)
	}

	for rows.Next() {
		skillVacancy := entity.SkillVacancy{}
		skill := entity.Skill{}
		vacancy := entity.Vacancy{}
		err = rows.Scan(
			&vacancy.ID,
			&skill.ID,
		)
		if err != nil {
			log.Printf("unable to scan row: %v\n", err)
			return nil, fmt.Errorf("unable to scan row: %v", err)
		}
		skillVacancy.S = skill
		skillVacancy.V = vacancy
		output = append(output, skillVacancy)
	}
	if rows.Err() != nil {
		log.Println(rows.Err())
		return nil, rows.Err()
	}

	return output, nil
}

func (r *Repo) DeleteSkillVacancy(ctx context.Context, input entity.SkillVacancyFilter) error {
	log.Println("skill-vacancy delete started")
	defer log.Println("skill-vacancy delete done")
	q := r.builder.Delete(entity.TableSkillVacancy)

	// Where
	// Skill part
	if len(input.SF.ID) > 0 {
		q = q.Where(squirrel.Eq{"id_skill": input.SF.ID})
	}
	// Vacancy part
	if len(input.VF.ID) > 0 {
		q = q.Where(squirrel.Eq{"id_vacancy": input.VF.ID})
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
		log.Printf("unable to delete skill-vacancy: %v\n", err)
		return fmt.Errorf("unable to delete skill-vacancy: %v", err)
	}

	return nil
}
