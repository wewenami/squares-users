package drivers

import (
	"context"

	"git.squares.dev/src/models"
)

type Base interface {
	Connect(ctx context.Context) error
	Ping(ctx context.Context) error
	Close(ctx context.Context) error

	UsersRepository() UsersRepository
}

type DataStore interface {
	Base
}

type UsersRepository interface {
	Create(ctx context.Context, u models.User) (string, error)
	GetById(ctx context.Context, id string) (models.User, error)
	List(ctx context.Context, paging models.Paging) ([]models.User, error)
	Update(ctx context.Context, id string, u models.User) (models.User, error)
	Delete(ctx context.Context, id string) error
}
