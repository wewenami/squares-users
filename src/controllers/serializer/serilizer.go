package serializer

import (
	"git.squares.dev/src/models"
	grpc "github.com/wewenami/grpc-proto/users"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func UserToProto(u models.User) *grpc.User {
	return &grpc.User{
		Id:         u.Id,
		FirstName:  u.FirstName,
		LastName:   u.LastName,
		Patronymic: u.Patronymic,
		Email:      u.Email,
		Phone:      u.Phone,
		IsEnabled:  u.IsEnabled,
		PhotoURL:   u.PhotoURL,
		CreatedAt:  timestamppb.New(u.CreatedAt),
		UpdateAt:   timestamppb.New(u.UpdateAt),
		DeletedAt:  timestamppb.New(u.DeletedAt),
	}
}

func UserListToProto(users []models.User) []*grpc.User {
	usersGrpc := make([]*grpc.User, 0, len(users))
	for _, u := range users {
		usersGrpc = append(usersGrpc, UserToProto(u))
	}

	return usersGrpc
}
