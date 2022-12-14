package main

import (
	"context"
	"net"
	"os/signal"
	"syscall"

	"github.com/GrabItYourself/giys-backend/lib/logger"
	"github.com/GrabItYourself/giys-backend/lib/postgres"
	"github.com/GrabItYourself/giys-backend/lib/postgres/repository"
	"github.com/GrabItYourself/giys-backend/lib/rabbitmq"
	"github.com/GrabItYourself/giys-backend/payment/internal/config"
	"github.com/GrabItYourself/giys-backend/payment/internal/server"
	"github.com/GrabItYourself/giys-backend/payment/pkg/paymentproto"
	"github.com/omise/omise-go"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
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

	// RabbitMQ Sender
	rabbitSender, err := rabbitmq.NewSender(conf.RabbitMQ.URL)
	if err != nil {
		panic(err)
	}
	defer rabbitSender.Close()

	// Initialize gRPC server
	grpcServer := grpc.NewServer()

	// Initialize PaymentService server
	omiseClient, e := omise.NewClient(conf.Omise.PublicKey, conf.Omise.SecretKey)
	if e != nil {
		logger.Fatal("Failed to initialize omise client: " + err.Error())
	}
	paymentServer, err := server.NewServer(omiseClient, repo, rabbitSender)
	if err != nil {
		logger.Fatal("Failed to initialize payment server: " + err.Error())
	}

	// Register PaymentService server
	paymentproto.RegisterPaymentServiceServer(grpcServer, paymentServer)

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
	if err = grpcServer.Serve(lis); err != nil {
		logger.Panic(errors.Wrap(err, "Failed to serve").Error())
	}
}
