package main

import (
	"context"
	"net"
	"os/signal"
	"syscall"

	"github.com/GrabItYourself/giys-backend/lib/logger"
	"github.com/GrabItYourself/giys-backend/lib/postgres"
	"github.com/GrabItYourself/giys-backend/lib/postgres/repository"
	"github.com/GrabItYourself/giys-backend/user/internal/config"
	"github.com/GrabItYourself/giys-backend/user/internal/server"
	"github.com/GrabItYourself/giys-backend/user/pkg/userproto"
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

	// Initialize Postgres connection
	pg, err := postgres.New(&conf.Postgres)
	if err != nil {
		logger.Fatal(errors.Wrap(err, "Can't initialize postgres").Error())
	}
	defer func() {
		if db, err := pg.DB(); err == nil {
			logger.Info("Closing database connection...")
			if err := db.Close(); err != nil {
				logger.Panic(errors.Wrap(err, "Failed to close database connection").Error())
			}
		} else {
			logger.Panic(errors.Wrap(err, "Can't close postgres").Error())
		}
	}()

	// Repository
	repo := repository.New(pg)

	// Initialize gRPC server
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	// Register UserService server
	userproto.RegisterUserServiceServer(grpcServer, server.NewServer(repo))

	// Serve
	lis, err := net.Listen("tcp", ":"+conf.Server.Port)
	if err != nil {
		logger.Panic(errors.Wrap(err, "Failed to listen").Error())
	}
	go func() {
		<-ctx.Done()
		logger.Info("Received shut down signal. Attempting graceful shutdown...")
		grpcServer.GracefulStop()
	}()

	logger.Info("Starting gRPC server on port " + conf.Server.Port)
	if err := grpcServer.Serve(lis); err != nil {
		logger.Panic(errors.Wrap(err, "Failed to serve").Error())
	}
}
