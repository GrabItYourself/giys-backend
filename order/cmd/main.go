package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"

	"github.com/GrabItYourself/giys-backend/lib/logger"
	"github.com/GrabItYourself/giys-backend/lib/postgres"
	"github.com/GrabItYourself/giys-backend/order/internal/config"
	"go.uber.org/zap"
)

func main() {
	// Load config
	cf := config.InitConfig()

	// Initialize logger
	logger.InitLogger(&cf.Log)

	// Initialize context
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	// Establish connection to PostgreSQL database
	pg, err := postgres.New(&cf.Postgres)
	if err != nil {
		logger.Fatal("Failed to initialize PostgreSQL connection", zap.Error(err))
	}

	fmt.Println(pg, ctx)
}
