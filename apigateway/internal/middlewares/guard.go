package middlewares

import (
	"github.com/GrabItYourself/giys-backend/auth/pkg/authproto"
	"github.com/GrabItYourself/giys-backend/auth/pkg/authutils"
	"github.com/GrabItYourself/giys-backend/lib/logger"
	"github.com/GrabItYourself/giys-backend/lib/postgres/models"
	"github.com/gofiber/fiber/v2"
)

const AccessTokenCookieName = "GIYS_ACCESS_TOKEN"
const RefreshTokenCookieName = "GIYS_REFRESH_TOKEN"

func NewAccessTokenGuard(authClient authproto.AuthClient) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		accessToken := c.Cookies(AccessTokenCookieName)
		if accessToken == "" {
			logger.Error("access token is empty")
			return fiber.NewError(fiber.StatusUnauthorized, "access token is empty")
		}
		verifyResp, err := authClient.VerifyAccessToken(c.Context(), &authproto.VerifyAccessTokenReq{
			AccessToken: accessToken,
		})
		// success
		if err == nil {
			logger.Debug("access token is valid, userId: " + verifyResp.UserId)
			c.Locals(authutils.IdentityKey, &authutils.Identity{
				UserId: verifyResp.UserId,
				Role:   models.RoleEnum(verifyResp.Role),
			})
			return c.Next()
		}

		// access token verification failed, try refreshing access token
		refreshToken := c.Cookies(RefreshTokenCookieName)
		if refreshToken == "" {
			logger.Error("refresh token is empty")
			return fiber.NewError(fiber.StatusUnauthorized, "access token verification failed and refresh token is empty")
		}
		refreshResp, err := authClient.RefreshAccessToken(c.Context(), &authproto.RefreshAccessTokenReq{
			RefreshToken: refreshToken,
		})
		// refresh token verification failed
		if err != nil {
			logger.Error("refresh token verification failed")
			return fiber.NewError(fiber.StatusUnauthorized, "access token verification failed and refresh token also verification failed")
		}

		// refresh token verification success, replace cookies
		c.Cookie(&fiber.Cookie{
			Name:     AccessTokenCookieName,
			Value:    refreshResp.AccessToken,
			HTTPOnly: true,
		})
		c.Cookie(&fiber.Cookie{
			Name:     RefreshTokenCookieName,
			Value:    refreshResp.RefreshToken,
			HTTPOnly: true,
		})
		c.Locals(authutils.IdentityKey, &authutils.Identity{
			UserId: refreshResp.UserId,
			Role:   models.RoleEnum(refreshResp.Role),
		})
		return c.Next()
	}
}
