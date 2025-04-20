package link

import (
	"context"
	"fmt"
	"log"

	"github.com/rashevskiivv/api/internal/entity"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

func (r *Repo) UpsertUserSkill(ctx context.Context, input entity.UserSkill) error {
	log.Println("user-skill upsert started")
	defer log.Println("user-skill upsert done")

	const q = `INSERT INTO user_skill ("id_user", "id_skill", "proficiency_level") 
VALUES (@id_user, @id_skill, @proficiency_level)
ON CONFLICT ON CONSTRAINT user_skill_pkey
	DO UPDATE SET proficiency_level	= EXCLUDED.proficiency_level;`

	args := pgx.NamedArgs{
		"id_user":           input.U.ID,
		"id_skill":          input.S.ID,
		"proficiency_level": input.ProficiencyLevel,
	}

	_, err := r.DB.Exec(ctx, q, args)
	if err != nil {
		log.Printf("unable to insert or update row: %v\n", err)
		return fmt.Errorf("unable to insert or update row: %v", err)
	}

	return nil
}

func (r *Repo) ReadUserSkill(ctx context.Context, input entity.UserSkillFilter) ([]entity.UserSkill, error) {
	log.Println("user-skill read started")
	defer log.Println("user-skill read done")
	var output []entity.UserSkill

	q := r.builder.Select(
		"id_user",
		"id_skill",
		"proficiency_level",
	).From(entity.TableUserSkill)

	// Where
	if len(input.SF.ID) > 0 {
		q = q.Where(squirrel.Eq{"id_skill": input.SF.ID})
	}
	if len(input.UF.ID) > 0 {
		q = q.Where(squirrel.Eq{"id_user": input.UF.ID})
	}
	if len(input.ProficiencyLevel) > 0 {
		q = q.Where(squirrel.Eq{"proficiency_level": input.ProficiencyLevel})
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
		userSkill := entity.UserSkill{}
		user := entity.User{}
		skill := entity.Skill{}
		err = rows.Scan(
			&user.ID,
			&skill.ID,
			&userSkill.ProficiencyLevel,
		)
		if err != nil {
			log.Printf("unable to scan row: %v\n", err)
			return nil, fmt.Errorf("unable to scan row: %v", err)
		}
		userSkill.S = skill
		userSkill.U = user
		output = append(output, userSkill)
	}
	if rows.Err() != nil {
		log.Println(rows.Err())
		return nil, rows.Err()
	}

	return output, nil
}

func (r *Repo) DeleteUserSkill(ctx context.Context, input entity.UserSkillFilter) error {
	log.Println("user-skill delete started")
	defer log.Println("user-skill delete done")
	q := r.builder.Delete(entity.TableUserSkill)

	// Where
	if len(input.SF.ID) > 0 {
		q = q.Where(squirrel.Eq{"id_skill": input.SF.ID})
	}
	if len(input.UF.ID) > 0 {
		q = q.Where(squirrel.Eq{"id_user": input.UF.ID})
	}
	if len(input.ProficiencyLevel) > 0 {
		q = q.Where(squirrel.Eq{"proficiency_level": input.ProficiencyLevel})
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
		log.Printf("unable to delete user-skill: %v\n", err)
		return fmt.Errorf("unable to delete user-skill: %v", err)
	}

	return nil
}
