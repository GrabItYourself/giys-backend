package v1router

import (
	"github.com/GrabItYourself/giys-backend/apigateway/internal/middlewares"
	"github.com/GrabItYourself/giys-backend/lib/postgres/models"
	"github.com/gofiber/fiber/v2"
)

type Body struct {
	UserId string          `json:"user_id"`
	Role   models.RoleEnum `json:"role"`
}

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
			return err
		}
		return c.JSON(user)
	})
}
