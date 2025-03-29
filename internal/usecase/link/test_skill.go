package link

import (
	"context"
	"log"
	"tax-api/internal/entity"
)

func (uc *UseCase) UpsertTS(ctx context.Context, input entity.TestSkill) error {
	log.Println("test-skill usecase upsert started")
	defer log.Println("test-skill usecase upsert done")

	err := input.Validate()
	if err != nil {
		log.Println(err)
		return err
	}

	err = uc.repo.UpsertTestSkill(ctx, input)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (uc *UseCase) DeleteTS(ctx context.Context, input entity.TestSkillFilter) error {
	log.Println("test-skill usecase delete started")
	defer log.Println("test-skill usecase delete done")

	err := input.Validate()
	if err != nil {
		log.Println(err)
		return err
	}

	err = uc.repo.DeleteTestSkill(ctx, input)
	if err != nil {
		log.Println(err)
	}
	return err
}
