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

	return nil

}

func (uc *UseCase) ReadUS(ctx context.Context, filter entity.UserSkillFilter) ([]entity.UserSkill, error) {
	err := filter.Validate()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return nil

}

func (uc *UseCase) DeleteUS(ctx context.Context, filter entity.UserSkillFilter) error {
	err := filter.Validate()
	if err != nil {
		log.Println(err)
		return err
	}

	return nil

}
