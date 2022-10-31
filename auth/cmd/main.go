package main

import (
	"context"
	"net"
	"os/signal"
	"syscall"

	"github.com/GrabItYourself/giys-backend/auth/internal/config"
	"github.com/GrabItYourself/giys-backend/auth/internal/repository"
	"github.com/GrabItYourself/giys-backend/auth/internal/server"
	"github.com/GrabItYourself/giys-backend/auth/pkg/authproto"
	"github.com/GrabItYourself/giys-backend/lib/logger"
	"github.com/GrabItYourself/giys-backend/lib/postgres"
	"github.com/GrabItYourself/giys-backend/lib/redis"
	"github.com/GrabItYourself/giys-backend/user/pkg/client"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
		logger.Panic(errors.Wrap(err, "Can't initialize postgres").Error())
	}
	defer func() {
		logger.Info("Closing database connection...")
		if db, err := pg.DB(); err != nil {
			logger.Panic(errors.Wrap(err, "Can't access postgres connection").Error())
		} else if err := db.Close(); err != nil {
			logger.Panic(errors.Wrap(err, "Can't close postgres connection").Error())
		}
	}()
	rdb, err := redis.New(ctx, &conf.Redis)
	if err != nil {
		logger.Panic(errors.Wrap(err, "Can't initialize redis").Error())
	}
	defer func() {
		logger.Info("Closing redis connection...")
		if err = rdb.Close(); err != nil {
			logger.Panic(errors.Wrap(err, "error during redis close").Error())
		}
	}()
	repo := repository.New(pg, rdb)

	// gRPC Clients
	userClient, conn, err := client.NewClient(ctx, conf.Grpc.User.Addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Panic(errors.Wrap(err, "Can't initialize user client").Error())
	}
	defer func() {
		logger.Info("Closing user client connection...")
		if err := conn.Close(); err != nil {
			logger.Panic(errors.Wrap(err, "Can't close user client connection").Error())
		}
	}()

	// Initialize gRPC server
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	// Register UserService server
	authproto.RegisterAuthServer(grpcServer, server.NewServer(repo, conf, userClient))

	// Serve
	lis, err := net.Listen("tcp", ":"+conf.Server.Port)
	if err != nil {
		logger.Panic(errors.Wrap(err, "Failed to listen").Error())
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
		logger.Panic(errors.Wrap(err, "Failed to serve").Error())
	}

}
