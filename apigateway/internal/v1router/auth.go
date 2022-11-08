package v1router

import (
	"github.com/GrabItYourself/giys-backend/apigateway/internal/middlewares"
	"github.com/GrabItYourself/giys-backend/apigateway/internal/v1router/types"
	"github.com/GrabItYourself/giys-backend/auth/pkg/authproto"
	"github.com/GrabItYourself/giys-backend/lib/logger"
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
		var clientTypeProto authproto.ClientType
		if clientType == "android" {
			clientTypeProto = authproto.ClientType_ANDROID
		} else if clientType == "ios" {
			clientTypeProto = authproto.ClientType_IOS
		} else {
			logger.Error("clientType is not valid")
			return fiber.NewError(fiber.StatusBadRequest, "clientType is invalid")
		}
		resp, err := r.Handler.HandleGoogleOAuthCallback(c.Context(), code, clientTypeProto)
		// TODO: infer error code from grpc error
		if err != nil {
			logger.Error(err.Error())
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
