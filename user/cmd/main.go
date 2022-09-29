package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	err := app.Listen(":3000")
	if err != nil {
		panic(errors.Wrap(err, "error during app listen"))
	}
}
