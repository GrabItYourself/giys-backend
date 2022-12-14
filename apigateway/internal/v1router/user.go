package v1router

import (
	"github.com/GrabItYourself/giys-backend/apigateway/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

func (r *Router) InitUserRoutes(basePath string) {
	f := r.Router.Group(basePath)

	f.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	// protect all paths below
	f.Use(middlewares.NewAccessTokenGuard(r.Handler.Grpc.Auth))

	f.Get("/me", func(c *fiber.Ctx) error {
		user, err := r.Handler.HandleUserMe(c)
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, err.Error())
		}
		return c.JSON(user)
	})

	f.Get("/orders", func(c *fiber.Ctx) error {
		order, err := r.Handler.HandleGetMyOrders(c)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		return c.JSON(order)
	})
}
