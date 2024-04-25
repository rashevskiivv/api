package repository

import "tax-api/internal/entity"

type UserRepo struct {
	// todo store db
}

type UserRepository interface {
	InsertUser(user entity.User) error
	ReadUsers(filter entity.Filter) ([]entity.User, error)
	UpdateUser(user entity.User, filter entity.Filter) error
	DeleteUser(filter entity.Filter) error

	InsertOperation(operation entity.Operation) error
	ReadOperations(filter entity.Filter) ([]entity.Operation, error)
	UpdateOperation(operation entity.Operation, filter entity.Filter) error
	DeleteOperation(filter entity.Filter) error
}

func (repo *UserRepo) InsertUser(user entity.User) error {

	return nil
}

func (repo *UserRepo) ReadUsers(filter entity.Filter) ([]entity.User, error) {
	var users []entity.User

	return users, nil
}

func (repo *UserRepo) UpdateUser(user entity.User, filter entity.Filter) error {

	return nil
}

func (repo *UserRepo) DeleteUser(filter entity.Filter) error {

	return nil
}

func (repo *UserRepo) InsertOperation(operation entity.Operation) error {

	return nil
}

func (repo *UserRepo) ReadOperations(filter entity.Filter) ([]entity.Operation, error) {
	var operations []entity.Operation

	return operations, nil
}

func (repo *UserRepo) UpdateOperation(operation entity.Operation, filter entity.Filter) error {

	return nil
}

func (repo *UserRepo) DeleteOperation(filter entity.Filter) error {

	return nil
}
