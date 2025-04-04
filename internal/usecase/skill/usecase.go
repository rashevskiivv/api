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
	err := input.Validate()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	output, err := uc.repo.Upsert(ctx, input)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (uc *UseCase) ReadSkills(ctx context.Context, input entity.SkillFilter) ([]entity.Skill, error) {
	err := input.Validate()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	output, err := uc.repo.Read(ctx, input)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (uc *UseCase) DeleteSkill(ctx context.Context, input entity.SkillFilter) error {
	err := input.Validate()
	if err != nil {
		log.Println(err)
		return err
	}

	return uc.repo.Delete(ctx, input)
}
