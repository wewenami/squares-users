package users_repo

import (
	"context"
	"database/sql"
	"time"

	"git.squares.dev/src/database/drivers"
	boiler "git.squares.dev/src/generated/sqlBoiler"
	"git.squares.dev/src/models"
	"github.com/google/uuid"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var (
	now = time.Now().UTC()
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) drivers.UsersRepository {
	return &repository{
		db: db,
	}
}

func (r repository) Create(ctx context.Context, u models.User) (string, error) {
	user := newUser(u)

	err := newUser(u).Insert(ctx, r.db, boil.Infer())
	if err != nil {
		return "", err
	}

	return user.ID, nil
}

func (r repository) GetById(ctx context.Context, id string) (models.User, error) {
	user, err := boiler.Users(boiler.UserWhere.ID.EQ(id)).One(ctx, r.db)
	if err != nil {
		return models.User{}, err
	}

	return toEntity(user), nil
}

func (r repository) List(ctx context.Context, paging models.Paging) ([]models.User, error) {
	usersOrm, err := boiler.Users().All(ctx, r.db)
	if err != nil {
		return nil, err
	}

	users := make([]models.User, 0, len(usersOrm))
	for _, u := range usersOrm {
		users = append(users, toEntity(u))
	}

	return users, nil
}

func (r repository) Update(ctx context.Context, id string, u models.User) (models.User, error) {
	user := updateUser(id, u)

	_, err := user.Update(ctx, r.db, boil.Infer())
	if err != nil {
		return models.User{}, err
	}

	return u, nil
}

func (r repository) Delete(ctx context.Context, id string) error {
	user, err := boiler.FindUser(ctx, r.db, id)
	if err != nil {
		return err
	}

	_, err = user.Delete(ctx, r.db)
	if err != nil {
		return err
	}

	return nil
}

func toEntity(u *boiler.User) models.User {
	if u == nil {
		return models.User{}
	}

	return models.User{
		Id:         u.ID,
		FirstName:  u.FirstName.String,
		LastName:   u.LastName.String,
		Patronymic: u.Patronymic.String,
		Phone:      u.Phone.String,
		Email:      u.Email.String,
		IsEnabled:  u.IsEnabled.Bool,
		PhotoURL:   u.PhotoURL.String,
		CreatedAt:  u.CreatedAt,
		UpdateAt:   u.UpdatedAt,
		DeletedAt:  u.DeletedAt.Time,
	}
}

func updateUser(id string, u models.User) *boiler.User {
	return &boiler.User{
		ID:         id,
		FirstName:  null.String{String: u.FirstName, Valid: true},
		LastName:   null.String{String: u.LastName, Valid: true},
		Patronymic: null.String{String: u.Patronymic, Valid: true},
		Phone:      null.String{String: u.Phone, Valid: true},
		Email:      null.String{String: u.Email, Valid: true},
		IsEnabled:  null.Bool{Bool: u.IsEnabled, Valid: true},
		PhotoURL:   null.String{String: u.PhotoURL, Valid: true},
		CreatedAt:  u.CreatedAt,
		UpdatedAt:  now,
		DeletedAt:  null.Time{Time: u.DeletedAt, Valid: true},
	}
}

func newUser(u models.User) *boiler.User {
	return &boiler.User{
		ID:         uuid.New().String(),
		FirstName:  null.String{String: u.FirstName, Valid: true},
		LastName:   null.String{String: u.LastName, Valid: true},
		Patronymic: null.String{String: u.Patronymic, Valid: true},
		Phone:      null.String{String: u.Phone, Valid: true},
		Email:      null.String{String: u.Email, Valid: true},
		IsEnabled:  null.Bool{Bool: true, Valid: true},
		PhotoURL:   null.String{String: u.PhotoURL, Valid: true},
		CreatedAt:  now,
		UpdatedAt:  now,
		DeletedAt:  null.Time{},
	}
}
