package link

import (
	"context"
	"log"
	"tax-api/internal/entity"
)

func (uc *UseCase) UpsertTS(ctx context.Context, input entity.TestSkill) error {
	err := input.Validate()
	if err != nil {
		log.Println(err)
		return err
	}

	return nil

}

func (uc *UseCase) ReadTS(ctx context.Context, filter entity.TestSkillFilter) ([]entity.TestSkill, error) {
	err := filter.Validate()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return nil

}

func (uc *UseCase) DeleteTS(ctx context.Context, filter entity.TestSkillFilter) error {
	err := filter.Validate()
	if err != nil {
		log.Println(err)
		return err
	}

	return nil

}
