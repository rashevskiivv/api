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

	return nil

}

func (uc *UseCase) ReadSV(ctx context.Context, filter entity.SkillVacancyFilter) ([]entity.SkillVacancy, error) {
	err := filter.Validate()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return nil

}

func (uc *UseCase) DeleteSV(ctx context.Context, filter entity.SkillVacancyFilter) error {
	err := filter.Validate()
	if err != nil {
		log.Println(err)
		return err
	}

	return nil

}
