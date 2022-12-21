package user

import (
	"context"
	"github.com/go-pg/pg/v10"
)

type RepositoryInterface interface {
	create(ctx context.Context, dto CreateUserDto) error
	getByEmail(email string) (*User, error)
}

type Repository struct {
	db *pg.DB
}

func NewUserRepository(db *pg.DB) RepositoryInterface {
	return &Repository{
		db,
	}
}

func (repo *Repository) create(ctx context.Context, user CreateUserDto) error {
	_, err := repo.db.Model(user).Context(ctx).Insert()

	if err != nil {
		return err
	}
	return nil
}

func (repo *Repository) getByEmail(email string) (*User, error) {
	user := &User{}
	err := repo.db.Model(user).Where("email = ?", email).Select()

	if err != nil {
		return user, err
	}
	return user, nil
}
