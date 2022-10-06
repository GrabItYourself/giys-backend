package v1router

import (
	"context"
	"encoding/json"

	"github.com/GrabItYourself/giys-backend/lib/postgres/models"
	"github.com/gofiber/fiber/v2"
)

type Body struct {
	UserId string          `json:"user_id"`
	Role   models.RoleEnum `json:"role"`
}

func (r *Router) InitUserRoutes(ctx context.Context, basePath string) {
	f := r.Router.Group(basePath)

	f.Get("/", func(c *fiber.Ctx) error {
		c.SendString("Hello, World 👋!")
		return nil
	})

	f.Get("/me", func(c *fiber.Ctx) error {
		body := &Body{}
		if err := json.Unmarshal(c.Body(), body); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		user, err := r.Handler.HandleUserMe(ctx, body.UserId, body.Role)
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, err.Error())
		}
		return c.SendString(user.String())
	})

	// ...
}
