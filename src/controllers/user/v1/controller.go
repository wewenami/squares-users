package v1

import (
	"context"

	"git.squares.dev/src/controllers/deserializer"
	"git.squares.dev/src/controllers/serializer"
	usersService "git.squares.dev/src/services/users"
	grpc "github.com/wewenami/grpc-proto/users"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Controller struct {
	usersService usersService.Service
	grpc.UnimplementedUsersServer
}

func NewController(usersService usersService.Service) *Controller {
	return &Controller{
		usersService: usersService,
	}
}

func (c Controller) Create(ctx context.Context, u *grpc.UserCreate) (*grpc.EntityId, error) {
	id, err := c.usersService.Create(ctx, deserializer.UserCreateFromProto(u))
	if err != nil {
		return nil, err
	}

	return &grpc.EntityId{Id: id}, nil

}

func (c Controller) GetById(ctx context.Context, u *grpc.EntityId) (*grpc.User, error) {
	if u == nil {
		return nil, nil
	}

	user, err := c.usersService.GetById(ctx, u.GetId())
	if err != nil {
		return nil, err
	}

	return serializer.UserToProto(user), nil
}

func (c Controller) List(ctx context.Context, p *grpc.Paging) (*grpc.UsersList, error) {
	if p == nil {
		return nil, nil
	}

	users, err := c.usersService.List(ctx, deserializer.PagingFromProto(p))
	if err != nil {
		return nil, err
	}

	return &grpc.UsersList{Users: serializer.UserListToProto(users)}, nil

}

func (c Controller) Update(ctx context.Context, u *grpc.UserUpdate) (*grpc.User, error) {
	if u == nil {
		return nil, nil
	}

	user, err := c.usersService.Update(ctx, u.GetId(), deserializer.UserUpdateFromProto(u))
	if err != nil {
		return nil, err
	}

	return serializer.UserToProto(user), nil

}

func (c Controller) Delete(ctx context.Context, u *grpc.EntityId) (*emptypb.Empty, error) {
	if u == nil {
		return nil, nil
	}

	err := c.usersService.Delete(ctx, u.GetId())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
