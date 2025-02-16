package skill

import (
	"context"
	"log"
	"tax-api/internal/entity"
	repositorySkill "tax-api/internal/repository/skill"
)

type UseCase struct {
	repo repositorySkill.Repository
}

func NewUseCase(repo repositorySkill.Repository) *UseCase {
	return &UseCase{repo: repo}
}

func (uc *UseCase) UpsertSkill(ctx context.Context, input entity.Skill) (*entity.Skill, error) {
	log.Println("skill usecase upsert started")

	output, err := uc.repo.Upsert(ctx, input)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("skill usecase upsert done")

	return output, nil
}

func (uc *UseCase) ReadSkills(ctx context.Context, input entity.SkillFilter) ([]entity.Skill, error) {
	log.Println("skill usecase read started")
	err := input.Validate()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	output, err := uc.repo.Read(ctx, input)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("skill usecase read done")

	return output, nil
}

func (uc *UseCase) DeleteSkill(ctx context.Context, input entity.SkillFilter) error {
	log.Println("skill usecase delete started")
	err := input.Validate()
	if err != nil {
		log.Println(err)
		return err
	}

	err = uc.repo.Delete(ctx, input)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("skill usecase delete done")

	return nil
}
