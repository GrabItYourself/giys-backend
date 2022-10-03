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
	config := config.InitConfig()

	// Initialize logger
	logger.InitLogger(&config.Log)

	// Initialize context
	context, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	// Establish connection to PostgreSQL database
	postgres, err := postgres.New(&config.Postgres)
	if err != nil {
		logger.Fatal("Failed to initialize PostgreSQL connection", zap.Error(err))
	}

	fmt.Println(postgres, context)
}
