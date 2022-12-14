package main

import (
	"context"
	"net"
	"os/signal"
	"syscall"

	"github.com/GrabItYourself/giys-backend/lib/logger"
	"github.com/GrabItYourself/giys-backend/lib/postgres"
	"github.com/GrabItYourself/giys-backend/lib/postgres/repository"
	"github.com/GrabItYourself/giys-backend/payment/pkg/client"
	"github.com/GrabItYourself/giys-backend/shop/internal/config"
	"github.com/GrabItYourself/giys-backend/shop/internal/server"
	"github.com/GrabItYourself/giys-backend/shop/pkg/shopproto"
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

	// Establish connection to PostgreSQL database
	pg, err := postgres.New(&conf.Postgres)
	if err != nil {
		logger.Fatal("Failed to initialize PostgreSQL connection: " + err.Error())
	}
	defer func() {
		logger.Info("Closing database connection...")
		if db, err := pg.DB(); err != nil {
			logger.Panic(errors.Wrap(err, "Can't access postgres connection").Error())
		} else if err := db.Close(); err != nil {
			logger.Panic(errors.Wrap(err, "Can't close postgres connection").Error())
		}
	}()
	repo := repository.New(pg)

	// gRPC Clients
	paymentClient, conn, err := client.NewClient(ctx, conf.Grpc.Payment.Addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Panic(errors.Wrap(err, "Can't initialize payment client").Error())
	}
	defer func() {
		logger.Info("Closing payment connection...")
		if err := conn.Close(); err != nil {
			logger.Panic(errors.Wrap(err, "Can't close payment connection").Error())
		}
	}()

	// Initialize gRPC server
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	// Register shop service implementation to the gRPC server
	shopproto.RegisterShopServiceServer(grpcServer, server.NewServer(repo, paymentClient))

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
