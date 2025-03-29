package link

import (
	"context"
	"log"
	"tax-api/internal/entity"
)

func (uc *UseCase) UpsertUS(ctx context.Context, input entity.UserSkill) error {
	log.Println("user-skill usecase upsert started")
	defer log.Println("user-skill usecase upsert done")

	err := input.Validate()
	if err != nil {
		log.Println(err)
		return err
	}

	err = uc.repo.UpsertUserSkill(ctx, input)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (uc *UseCase) ReadUS(ctx context.Context, input entity.UserSkillFilter) ([]entity.UserSkill, error) {
	log.Println("user-skill usecase read started")
	defer log.Println("user-skill usecase read done")

	err := input.Validate()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	output, err := uc.repo.ReadUserSkill(ctx, input)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return output, nil
}

func (uc *UseCase) DeleteUS(ctx context.Context, input entity.UserSkillFilter) error {
	log.Println("user-skill usecase delete started")
	defer log.Println("user-skill usecase delete done")

	err := input.Validate()
	if err != nil {
		log.Println(err)
		return err
	}

	err = uc.repo.DeleteUserSkill(ctx, input)
	if err != nil {
		log.Println(err)
	}
	return err
}
