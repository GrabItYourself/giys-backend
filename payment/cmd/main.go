package main

import (
	"net"

	"github.com/GrabItYourself/giys-backend/lib/logger"
	pb "github.com/GrabItYourself/giys-backend/lib/proto/payment"
	"github.com/GrabItYourself/giys-backend/payment/config"
	"github.com/GrabItYourself/giys-backend/payment/pkg/server"
	"google.golang.org/grpc"
)

func main() {
	conf := config.InitConfig()
	logger.InitLogger(&conf.Log)

	s := grpc.NewServer()
    lis, err := net.Listen("tcp", ":" + conf.Server.Port)
    if err != nil {
		logger.Fatal("Failed to listen: " + err.Error())
    }

	pb.RegisterPaymentServer(s, server.NewServer())

	err = s.Serve(lis)
    if err != nil {
		logger.Fatal("Failed to serve: " + err.Error())
    }
}
