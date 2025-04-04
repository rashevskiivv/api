package link

import (
	"context"
	"log"
	"tax-api/internal/entity"
)

func (uc *UseCase) UpsertUS(ctx context.Context, input entity.UserSkill) error {
	err := input.Validate()
	if err != nil {
		log.Println(err)
		return err
	}

	return uc.repo.UpsertUserSkill(ctx, input)
}

func (uc *UseCase) ReadUS(ctx context.Context, input entity.UserSkillFilter) ([]entity.UserSkill, error) {
	err := input.Validate()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	output, err := uc.repo.ReadUserSkill(ctx, input)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (uc *UseCase) DeleteUS(ctx context.Context, input entity.UserSkillFilter) error {
	err := input.Validate()
	if err != nil {
		log.Println(err)
		return err
	}

	return uc.repo.DeleteUserSkill(ctx, input)
}
