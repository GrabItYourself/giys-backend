package main

import (
	"net"

	"github.com/GrabItYourself/giys-backend/lib/logger"
	"github.com/GrabItYourself/giys-backend/lib/postgres"
	"github.com/GrabItYourself/giys-backend/user/internal/config"
	"github.com/GrabItYourself/giys-backend/user/internal/libproto"
	"github.com/GrabItYourself/giys-backend/user/internal/repository"
	"github.com/GrabItYourself/giys-backend/user/internal/server"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Config
	conf := config.InitConfig()

	// Logger
	logger.InitLogger(&conf.Log)

	// Repository
	pg, err := postgres.New(&conf.Postgres)
	if err != nil {
		logger.Fatal(errors.Wrap(err, "Can't initialize postgres").Error())
	}
	repo := repository.New(pg)

	// Initialize gRPC server
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	// Register UserService server
	libproto.RegisterUserServiceServer(grpcServer, server.NewServer(repo))

	// Serve
	lis, err := net.Listen("tcp", ":"+conf.Server.Port)
	if err != nil {
		logger.Fatal(errors.Wrap(err, "Failed to listen").Error())
	}
	logger.Info("Starting gRPC server on port " + conf.Server.Port)
	err = grpcServer.Serve(lis)
	if err != nil {
		logger.Fatal(errors.Wrap(err, "Failed to serve").Error())
	}

}
