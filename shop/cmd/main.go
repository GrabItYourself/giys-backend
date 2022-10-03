package main

import (
	"fmt"

	"github.com/GrabItYourself/giys-backend/lib/logger"
	"github.com/GrabItYourself/giys-backend/lib/postgres"
	"github.com/GrabItYourself/giys-backend/shop/internal/config"
)

func main() {
	// Config
	conf := config.InitConfig()

	// Logger
	logger.InitLogger(&conf.Log)

	// Establish connection to PostgreSQL database
	postgres, err := postgres.New(&conf.Postgres)
	if err != nil {
		logger.Fatal("Failed to initialize PostgreSQL connection: " + err.Error())
	}

	fmt.Println(postgres)
}
