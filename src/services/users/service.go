package users

import (
	"context"

	"git.squares.dev/src/database/drivers"
	"git.squares.dev/src/models"
)

type Service interface {
	Create(ctx context.Context, u models.User) (string, error)
	GetById(ctx context.Context, id string) (models.User, error)
	List(ctx context.Context, paging models.Paging) ([]models.User, error)
	Update(ctx context.Context, id string, u models.User) (models.User, error)
	Delete(ctx context.Context, id string) error
}

type service struct {
	usersRepo drivers.UsersRepository
}

func NewService(usersRepo drivers.UsersRepository) Service {
	return &service{
		usersRepo: usersRepo,
	}
}

func (s service) Create(ctx context.Context, u models.User) (string, error) {
	return s.usersRepo.Create(ctx, u)
}

func (s service) GetById(ctx context.Context, id string) (models.User, error) {
	return s.usersRepo.GetById(ctx, id)
}

func (s service) List(ctx context.Context, paging models.Paging) ([]models.User, error) {
	return s.usersRepo.List(ctx, paging)
}

func (s service) Update(ctx context.Context, id string, u models.User) (models.User, error) {
	return s.usersRepo.Update(ctx, id, u)
}

func (s service) Delete(ctx context.Context, id string) error {
	return s.usersRepo.Delete(ctx, id)
}
