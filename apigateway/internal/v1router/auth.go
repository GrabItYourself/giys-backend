package v1router

import (
	"context"
	"encoding/base64"
	"math/rand"
	"time"

	"github.com/GrabItYourself/giys-backend/apigateway/internal/v1router/types"
	"github.com/gofiber/fiber/v2"
)

func (r *Router) InitAuthRoutes(ctx context.Context, basePath string) {
	f := r.Router.Group(basePath)

	f.Get("/google", func(c *fiber.Ctx) error {
		cookie, state := generateOAuthStateCookie()
		c.Cookie(cookie)
		url := r.Handler.OAuthConfig.AuthCodeURL(state)
		return c.Redirect(url)
	})

	f.Get("/google/callback", func(c *fiber.Ctx) error {
		state := c.Query("state")
		cookie := c.Cookies(OauthCookieName)
		if state != cookie {
			return fiber.NewError(fiber.StatusBadRequest, "Invalid OAuth state")
		}

		code := c.Query("code")
		resp, err := r.Handler.HandleGoogleOAuthCallback(ctx, code)
		// TODO: handle grpc error better
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.JSON(types.GoogleOAuthCallbackResp{
			AccessToken:  resp.AccessToken,
			RefreshToken: resp.RefreshToken,
		})
	})
}

const OauthCookieName = "oauth-state"

func generateOAuthStateCookie() (*fiber.Cookie, string) {
	var expiration = time.Now().Add(365 * 24 * time.Hour)
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)

	return &fiber.Cookie{
		Name:     OauthCookieName,
		Value:    state,
		Expires:  expiration,
		HTTPOnly: true,
	}, state
}
