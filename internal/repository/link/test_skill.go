package link

import (
	"context"
	"fmt"
	"log"
	"tax-api/internal/entity"

	"github.com/jackc/pgx/v5"
)

func (r *Repo) UpsertTestSkill(ctx context.Context, input entity.TestSkill) error {
	log.Println("test-skill upsert started")
	defer log.Println("test-skill upsert done")

	const q = `UPDATE test
SET "id_skill"=@id_skill
WHERE "id"=@id_test;`

	args := pgx.NamedArgs{
		"id_skill": input.S.ID,
		"id_test":  input.T.ID,
	}

	_, err := r.DB.Exec(ctx, q, args)
	if err != nil {
		log.Printf("unable to insert or update row: %v\n", err)
		return fmt.Errorf("unable to insert or update row: %v", err)
	}

	return nil
}

func (r *Repo) DeleteTestSkill(ctx context.Context, input entity.TestSkillFilter) error {
	log.Println("test-skill delete started")
	defer log.Println("test-skill delete done")

	const q = `UPDATE test SET "id_skill"=null WHERE "id"=@id_test;`

	args := pgx.NamedArgs{
		"id_test": input.TF.ID,
	}

	_, err := r.DB.Exec(ctx, q, args)
	if err != nil {
		log.Printf("unable to insert or update row: %v\n", err)
		return fmt.Errorf("unable to insert or update row: %v", err)
	}

	return nil
}
