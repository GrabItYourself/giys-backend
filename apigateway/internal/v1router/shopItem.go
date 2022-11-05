package v1router

import (
	"strconv"

	"github.com/GrabItYourself/giys-backend/apigateway/internal/middlewares"
	"github.com/GrabItYourself/giys-backend/shop/pkg/shopproto"
	"github.com/gofiber/fiber/v2"
)

func (r *Router) InitShopItemRoutes(basePath string) {
	f := r.Router.Group(basePath)

	// protect all paths below
	f.Use(middlewares.NewAccessTokenGuard(r.Handler.Grpc.Auth))

	f.Get("/:shopItemId", func(c *fiber.Ctx) error {
		shopIdStr, shopItemIdStr := c.Params("shopId"), c.Params("shopItemId")
		shopId, err := strconv.Atoi(shopIdStr)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "shopId is not a number")
		}
		shopItemId, err := strconv.Atoi(shopItemIdStr)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "shopItemId is not a number")
		}

		shopItem, err := r.Handler.HandleGetShopItem(c, int32(shopId), int32(shopItemId))
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		return c.JSON(shopItem)
	})

	f.Get("/", func(c *fiber.Ctx) error {
		shopIdStr := c.Params("shopId")
		shopId, err := strconv.Atoi(shopIdStr)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "shopId is not a number")
		}

		shopItems, err := r.Handler.HandleGetAllShopItems(c, int32(shopId))
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		return c.JSON(shopItems)
	})

	f.Post("/", func(c *fiber.Ctx) error {
		reqBody := new(shopproto.CreateShopItemRequest)
		if err := c.BodyParser(reqBody); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "shopItem is not valid json")
		}
		shopItem, err := r.Handler.HandleCreateShopItem(c, reqBody)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		return c.JSON(shopItem)
	})

	f.Put("/:shopItemId", func(c *fiber.Ctx) error {
		shopIdStr, shopItemIdStr := c.Params("shopId"), c.Params("shopItemId")
		shopId, err := strconv.Atoi(shopIdStr)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "shopId is not a number")
		}
		shopItemId, err := strconv.Atoi(shopItemIdStr)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "shopItemId is not a number")
		}

		reqBody := new(shopproto.EditShopItemRequest)
		if err := c.BodyParser(reqBody); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "shopItem is not valid json")
		}

		shopItem, err := r.Handler.HandleEditShopItem(c, int32(shopId), int32(shopItemId), reqBody.EditedItem)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		return c.JSON(shopItem)
	})

	f.Delete("/:shopItemId", func(c *fiber.Ctx) error {
		shopIdStr, shopItemIdStr := c.Params("shopId"), c.Params("shopItemId")
		shopId, err := strconv.Atoi(shopIdStr)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "shopId is not a number")
		}
		shopItemId, err := strconv.Atoi(shopItemIdStr)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "shopItemId is not a number")
		}

		res, err := r.Handler.HandleDeleteShopItem(c, int32(shopId), int32(shopItemId))
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		return c.JSON(res)
	})
}