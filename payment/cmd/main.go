package main

import (
	"net"

	"github.com/GrabItYourself/giys-backend/lib/logger"
	"github.com/GrabItYourself/giys-backend/lib/postgres"
	"github.com/GrabItYourself/giys-backend/payment/internal/config"
	"github.com/GrabItYourself/giys-backend/payment/internal/libproto"
	"github.com/GrabItYourself/giys-backend/payment/internal/repository"
	"github.com/GrabItYourself/giys-backend/payment/internal/server"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

func main() {
	conf := config.InitConfig()
	logger.InitLogger(&conf.Log)

	pg, err := postgres.New(&conf.Postgres)
	if err != nil {
		logger.Fatal(errors.Wrap(err, "Can't initialize postgres").Error())
	}
	repo := repository.New(pg)

	s := grpc.NewServer()
	lis, err := net.Listen("tcp", ":"+conf.Server.Port)
	if err != nil {
		logger.Fatal("Failed to listen: " + err.Error())
	}

	paymentServer, err := server.NewServer(&conf.Omise, repo)
	if err != nil {
		logger.Fatal("Failed to initialize payment server: " + err.Error())
	}

	libproto.RegisterPaymentServiceServer(s, paymentServer)

	err = s.Serve(lis)
	if err != nil {
		logger.Fatal("Failed to serve: " + err.Error())
	}
}
