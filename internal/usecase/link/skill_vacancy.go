package link

import (
	"context"
	"log"
	"tax-api/internal/entity"
)

func (uc *UseCase) UpsertSV(ctx context.Context, input entity.SkillVacancy) error {
	err := input.Validate()
	if err != nil {
		log.Println(err)
		return err
	}

	return uc.repo.UpsertSkillVacancy(ctx, input)
}

func (uc *UseCase) ReadSV(ctx context.Context, input entity.SkillVacancyFilter) ([]entity.SkillVacancy, error) {
	err := input.Validate()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	output, err := uc.repo.ReadSkillVacancy(ctx, input)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (uc *UseCase) DeleteSV(ctx context.Context, input entity.SkillVacancyFilter) error {
	err := input.Validate()
	if err != nil {
		log.Println(err)
		return err
	}

	return uc.repo.DeleteSkillVacancy(ctx, input)
}
