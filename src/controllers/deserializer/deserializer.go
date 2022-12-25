package deserializer

import (
	"git.squares.dev/src/models"
	grpc "github.com/wewenami/grpc-proto/users"
)

func UserFromProto(u *grpc.User) models.User {
	if u == nil {
		return models.User{}
	}

	return models.User{
		Id:         u.GetId(),
		FirstName:  u.GetFirstName(),
		LastName:   u.GetLastName(),
		Patronymic: u.GetPatronymic(),
		Phone:      u.GetPhone(),
		Email:      u.GetEmail(),
		IsEnabled:  u.GetIsEnabled(),
		PhotoURL:   u.GetPhotoURL(),
		CreatedAt:  u.GetCreatedAt().AsTime(),
		UpdateAt:   u.GetUpdateAt().AsTime(),
		DeletedAt:  u.GetDeletedAt().AsTime(),
	}
}

func UserCreateFromProto(u *grpc.UserCreate) models.User {
	if u == nil {
		return models.User{}
	}

	return models.User{
		Id:         u.GetId(),
		FirstName:  u.GetFirstName(),
		LastName:   u.GetLastName(),
		Patronymic: u.GetPatronymic(),
		Phone:      u.GetPhone(),
		Email:      u.GetEmail(),
		IsEnabled:  u.GetIsEnabled(),
		PhotoURL:   u.GetPhotoURL(),
	}
}

func PagingFromProto(p *grpc.Paging) models.Paging {
	if p == nil {
		return models.Paging{}
	}

	return models.Paging{
		Page:  p.GetPage(),
		Limit: p.GetPage(),
	}
}

func UserUpdateFromProto(u *grpc.UserUpdate) models.User {
	if u == nil {
		return models.User{}
	}

	var user models.User

	if u.FirstName != nil {
		user.FirstName = u.GetFirstName()
	}

	if u.LastName != nil {
		user.LastName = u.GetLastName()
	}

	if u.Patronymic != nil {
		user.Patronymic = u.GetPatronymic()
	}

	if u.Email != nil {
		user.Email = u.GetEmail()
	}

	if u.Phone != nil {
		user.Phone = u.GetPhone()
	}

	if u.IsEnabled != nil {
		user.IsEnabled = u.GetIsEnabled()
	}

	if u.PhotoURL != nil {
		user.PhotoURL = u.GetPhotoURL()
	}

	return user
}
