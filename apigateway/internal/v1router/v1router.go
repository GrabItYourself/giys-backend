package v1router

import (
	"context"

	v1handler "github.com/GrabItYourself/giys-backend/apigateway/internal/v1router/handler"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	Ctx     context.Context
	Router  fiber.Router
	Handler *v1handler.Handler
}

func NewRouter(ctx context.Context, r fiber.Router, h *v1handler.Handler) *Router {
	return &Router{
		Ctx:     ctx,
		Router:  r,
		Handler: h,
	}
}
