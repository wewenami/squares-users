package main

import (
	"context"
	"log"
	"net"
	"os/signal"
	"syscall"

	"git.squares.dev/src/config"
	v1 "git.squares.dev/src/controllers/user/v1"
	"git.squares.dev/src/database"
	"git.squares.dev/src/database/drivers"
	"git.squares.dev/src/services/users"
	users2 "github.com/wewenami/grpc-proto/users"
	"google.golang.org/grpc"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	var opts config.Config

	if err := config.Parse(&opts); err != nil {
		log.Fatal("[ERROR]")
	}

	ds, err := database.New(drivers.Config{
		DBName: opts.DBName,
		DSName: opts.DSName,
		URL:    opts.DBURL,
	})
	defer ds.Close(ctx)

	if err = ds.Connect(ctx); err == nil {
		log.Fatal("[ERROR]")
	}

	if err = ds.Ping(ctx); err != nil {
		log.Fatal("[ERROR]")
	}

	usersService := users.NewService(ds.UsersRepository())
	usersController := v1.NewController(usersService)

	srv := grpc.NewServer()
	srv.RegisterService(&users2.Users_ServiceDesc, &usersController)

	listener, err := net.Listen("tcp", opts.GrpcListenPort)

	err = srv.Serve(listener)
}
