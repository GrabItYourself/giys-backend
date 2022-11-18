package v1router

import (
	"strconv"

	"github.com/GrabItYourself/giys-backend/apigateway/internal/middlewares"
	"github.com/GrabItYourself/giys-backend/apigateway/internal/v1router/types"
	"github.com/GrabItYourself/giys-backend/lib/logger"
	"github.com/GrabItYourself/giys-backend/shop/pkg/shopproto"
	"github.com/gofiber/fiber/v2"
)

func (r *Router) InitShopRoutes(basePath string) {
	f := r.Router.Group(basePath)

	// protect all paths below
	f.Use(middlewares.NewAccessTokenGuard(r.Handler.Grpc.Auth))

	f.Get("/:shopId", func(c *fiber.Ctx) error {
		shopIdStr := c.Params("shopId")
		shopId, err := strconv.Atoi(shopIdStr)
		if err != nil {
			logger.Error(err.Error())
			return fiber.NewError(fiber.StatusBadRequest, "shopId is not a number")
		}
		shop, err := r.Handler.HandleGetShop(c, int32(shopId))
		if err != nil {
			logger.Error(err.Error())
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		return c.JSON(shop)
	})

	f.Get("/", func(c *fiber.Ctx) error {
		shops, err := r.Handler.HandleGetAllShops(c)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		return c.JSON(shops)
	})

	f.Post("/", func(c *fiber.Ctx) error {
		var reqBody types.CreateShopWithBankAccountRequest
		if err := c.BodyParser(&reqBody); err != nil {
			logger.Error(err.Error())
			return fiber.NewError(fiber.StatusBadRequest, "shop is not valid json")
		}
		shop, err := r.Handler.HandleCreateShop(c, &reqBody)
		if err != nil {
			logger.Error(err.Error())
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		return c.JSON(shop)
	})

	f.Put("/:shopId", func(c *fiber.Ctx) error {
		shopIdStr := c.Params("shopId")
		shopId, err := strconv.Atoi(shopIdStr)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "shopId is not a number")
		}

		var reqBody shopproto.EditShopRequest
		if err := c.BodyParser(&reqBody); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "shop is not valid json")
		}

		shop, err := r.Handler.HandleEditShop(c, int32(shopId), reqBody.EditedShop)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		return c.JSON(shop)
	})

	f.Put("/:shopId/owners", func(c *fiber.Ctx) error {
		shopIdStr := c.Params("shopId")
		shopId, err := strconv.Atoi(shopIdStr)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "shopId is not a number")
		}

		var reqBody shopproto.EditShopOwnersRequest
		if err := c.BodyParser(&reqBody); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "shop is not valid json")
		}

		shop, err := r.Handler.HandleEditShopOwners(c, int32(shopId), reqBody.OwnerEmails)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		return c.JSON(shop)
	})

	f.Delete("/:shopId", func(c *fiber.Ctx) error {
		shopIdStr := c.Params("shopId")
		shopId, err := strconv.Atoi(shopIdStr)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "shopId is not a number")
		}

		res, err := r.Handler.HandleDeleteShop(c, int32(shopId))
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		return c.JSON(res)
	})
}
