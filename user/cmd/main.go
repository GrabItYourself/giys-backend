package main

import (
	"github.com/GrabItYourself/giys-backend/lib/logger"
	"github.com/GrabItYourself/giys-backend/user/config"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

func main() {
	conf := config.InitConfig()
	err := logger.InitLogger(&conf.Log)
	if err != nil {
		logger.Fatal(errors.Wrap(err, "can't initialize logger").Error())
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	err = app.Listen(":" + conf.Server.Port)
	if err != nil {
		logger.Fatal(errors.Wrap(err, "error during app listen").Error())
	}
}
