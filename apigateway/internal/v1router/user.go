package v1router

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

func (r *Router) InitUserRoute(ctx context.Context, basePath string) {
	f := r.Router.Group(basePath)

	f.Get("/me", func(c *fiber.Ctx) error {
		user, err := r.Handler.HandleUserMe(ctx)
		if err != nil {
			return errors.WithStack(err)
		}
		return c.SendString(user.String())
	})

	// ...
}
