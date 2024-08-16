package repository

import (
	"tax-api/internal/entity"

	"github.com/Masterminds/squirrel"
)

type OperationRepo struct {
	Postgres
	builder squirrel.StatementBuilderType
}

type OperationRepository interface {
	InsertOperation(operation entity.Operation) error
	ReadOperations(filter entity.Filter) ([]entity.Operation, error)
	UpdateOperation(operation entity.Operation, filter entity.Filter) error
	DeleteOperation(filter entity.Filter) error
}

func (repo *OperationRepo) InsertOperation(operation entity.Operation) error {
	// todo implement
	return nil
}

func (repo *OperationRepo) ReadOperations(filter entity.Filter) ([]entity.Operation, error) {
	var operations []entity.Operation
	// todo implement
	return operations, nil
}

func (repo *OperationRepo) UpdateOperation(operation entity.Operation, filter entity.Filter) error {
	// todo implement
	return nil
}

func (repo *OperationRepo) DeleteOperation(filter entity.Filter) error {
	// todo implement
	return nil
}
