package link

import (
	"context"
	"log"

	"github.com/rashevskiivv/api/internal/entity"
)

func (uc *UseCase) UpsertTS(ctx context.Context, input entity.TestSkill) error {
	err := input.Validate()
	if err != nil {
		log.Println(err)
		return err
	}

	return uc.repo.UpsertTestSkill(ctx, input)
}

func (uc *UseCase) DeleteTS(ctx context.Context, input entity.TestSkillFilter) error {
	err := input.Validate()
	if err != nil {
		log.Println(err)
		return err
	}

	return uc.repo.DeleteTestSkill(ctx, input)
}
