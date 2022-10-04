package main

import (
	"context"
	"net"
	"os/signal"
	"syscall"

	"github.com/GrabItYourself/giys-backend/auth/internal/config"
	"github.com/GrabItYourself/giys-backend/auth/internal/libproto"
	"github.com/GrabItYourself/giys-backend/auth/internal/repository"
	"github.com/GrabItYourself/giys-backend/auth/internal/server"
	"github.com/GrabItYourself/giys-backend/lib/logger"
	"github.com/GrabItYourself/giys-backend/lib/postgres"
	"github.com/GrabItYourself/giys-backend/lib/redis"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Context
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	// Config
	conf := config.InitConfig()

	// Logger
	logger.InitLogger(&conf.Log)

	// Repository
	pg, err := postgres.New(&conf.Postgres)
	if err != nil {
		logger.Fatal(errors.Wrap(err, "Can't initialize postgres").Error())
	}
	rdb, err := redis.New(ctx, &conf.Redis)
	repo := repository.New(pg, rdb)

	// Initialize gRPC server
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	// Register UserService server
	libproto.RegisterAuthServer(grpcServer, server.NewServer(repo, &conf.OAuth))

	// Serve
	lis, err := net.Listen("tcp", ":"+conf.Server.Port)
	if err != nil {
		logger.Fatal(errors.Wrap(err, "Failed to listen").Error())
	}
	logger.Info("Starting gRPC server on port " + conf.Server.Port)
	go func() {
		<-ctx.Done()
		cancel()
		logger.Info("Received shut down signal. Attempting graceful shutdown...")
		grpcServer.GracefulStop()
	}()
	err = grpcServer.Serve(lis)
	if err != nil {
		logger.Fatal(errors.Wrap(err, "Failed to serve").Error())
	}

}
