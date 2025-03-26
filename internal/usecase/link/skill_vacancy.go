package link

import (
	"context"
	"log"
	"tax-api/internal/entity"
)

func (uc *UseCase) UpsertSV(ctx context.Context, input entity.SkillVacancy) error {
	log.Println("skill-vacancy usecase upsert started")
	defer log.Println("skill-vacancy usecase upsert done")

	err := input.Validate()
	if err != nil {
		log.Println(err)
		return err
	}

	err = uc.repo.UpsertSkillVacancy(ctx, input)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (uc *UseCase) ReadSV(ctx context.Context, input entity.SkillVacancyFilter) ([]entity.SkillVacancy, error) {
	log.Println("skill-vacancy usecase read started")
	defer log.Println("skill-vacancy usecase read done")

	err := input.Validate()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	output, err := uc.repo.ReadSkillVacancy(ctx, input)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return output, nil
}

func (uc *UseCase) DeleteSV(ctx context.Context, input entity.SkillVacancyFilter) error {
	log.Println("skill-vacancy usecase delete started")
	defer log.Println("skill-vacancy usecase delete done")

	err := input.Validate()
	if err != nil {
		log.Println(err)
		return err
	}

	err = uc.repo.DeleteSkillVacancy(ctx, input)
	if err != nil {
		log.Println(err)
	}
	return err

}
