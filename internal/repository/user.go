package repository

import (
	"context"
	"fmt"
	env "tax-api/internal"
	"tax-api/internal/entity"

	"github.com/jackc/pgx/v5"
)

type UserRepo struct {
	Postgres
}

type UserRepository interface {
	InsertUser(ctx context.Context, user entity.User) error
	ReadUsers(filter entity.Filter) ([]entity.User, error)
	UpdateUser(user entity.User, filter entity.Filter) error
	DeleteUser(filter entity.Filter) error

	InsertOperation(operation entity.Operation) error
	ReadOperations(filter entity.Filter) ([]entity.Operation, error)
	UpdateOperation(operation entity.Operation, filter entity.Filter) error
	DeleteOperation(filter entity.Filter) error
}

func NewUserRepo(ctx context.Context) UserRepo {
	return UserRepo{NewPG(ctx, env.GetDBUrlEnv())}
}

func (repo *UserRepo) InsertUser(ctx context.Context, user entity.User) error {
	query := `INSERT INTO public."User" ("Name", "INN", "Email", "Password") VALUES (@name, @inn, @email, @password);`
	args := pgx.NamedArgs{
		"name":     user.Name,
		"inn":      user.INN,
		"email":    user.Email,
		"password": user.Password,
	}
	_, err := repo.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}
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
