package main

import (
	"github.com/GrabItYourself/giys-backend/lib/logger"
	"github.com/GrabItYourself/giys-backend/lib/postgres"
	"github.com/GrabItYourself/giys-backend/user/config"
	"github.com/GrabItYourself/giys-backend/user/pkg/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

func main() {
	conf := config.InitConfig()
	logger.InitLogger(&conf.Log)

	pg, err := postgres.New(&conf.Postgres)
	if err != nil {
		logger.Fatal(errors.Wrap(err, "Can't initialize postgres").Error())
	}
	repo := repository.New(pg)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	err = app.Listen(":" + conf.Server.Port)
	if err != nil {
		logger.Fatal(errors.Wrap(err, "error during app listen").Error())
	}
}
