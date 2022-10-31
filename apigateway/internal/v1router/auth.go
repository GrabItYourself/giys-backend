package v1router

import (
	"github.com/GrabItYourself/giys-backend/apigateway/internal/middlewares"
	"github.com/GrabItYourself/giys-backend/apigateway/internal/v1router/types"
	"github.com/GrabItYourself/giys-backend/auth/pkg/authproto"
	"github.com/gofiber/fiber/v2"
)

func (r *Router) InitAuthRoutes(basePath string) {
	f := r.Router.Group(basePath)

	f.Post("/google/verify", func(c *fiber.Ctx) error {
		code := c.Query("code")
		if code == "" {
			return fiber.NewError(fiber.StatusBadRequest, "code is required")
		}
		clientType := c.Query("clientType")
		if clientType == "" {
			return fiber.NewError(fiber.StatusBadRequest, "clientType is required")
		}
		clientTypeEnum, ok := authproto.ClientType_value[clientType]
		if !ok {
			return fiber.NewError(fiber.StatusBadRequest, "invalid client")
		}
		resp, err := r.Handler.HandleGoogleOAuthCallback(c.Context(), code, authproto.ClientType(clientTypeEnum))
		// TODO: infer error code from grpc error
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		c.Cookie(&fiber.Cookie{
			Name:     middlewares.AccessTokenCookieName,
			Value:    resp.AccessToken,
			HTTPOnly: true,
		})
		c.Cookie(&fiber.Cookie{
			Name:     middlewares.RefreshTokenCookieName,
			Value:    resp.RefreshToken,
			HTTPOnly: true,
		})

		return c.JSON(types.GoogleOAuthCallbackResp{
			AccessToken:  resp.AccessToken,
			RefreshToken: resp.RefreshToken,
		})
	})
}
