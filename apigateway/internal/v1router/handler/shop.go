package v1handler

import (
	"github.com/GrabItYourself/giys-backend/auth/pkg/authutils"
	"github.com/GrabItYourself/giys-backend/shop/pkg/shopproto"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

// ----- shop -----

func (h *Handler) HandleGetShop(c *fiber.Ctx, shopId int32) (*shopproto.ShopResponse, error) {
	identity, ok := c.Locals(authutils.IdentityKey).(*authutils.Identity)
	if !ok {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "identity not found in context")
	}
	ctx, err := authutils.EmbedIdentityToContext(c.Context(), identity)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, errors.Wrap(err, "can't embed identity to grpc context").Error())
	}

	res, err := h.Grpc.Shop.GetShop(ctx, &shopproto.GetShopRequest{
		Id: shopId,
	})
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, errors.Wrap(err, "Failed to request GRPC shop").Error())
	}
	return res, nil
}

func (h *Handler) HandleCreateShop(c *fiber.Ctx, shop *shopproto.CreateShopRequest) (*shopproto.ShopResponse, error) {
	identity, ok := c.Locals(authutils.IdentityKey).(*authutils.Identity)
	if !ok {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "identity not found in context")
	}
	ctx, err := authutils.EmbedIdentityToContext(c.Context(), identity)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, errors.Wrap(err, "can't embed identity to grpc context").Error())
	}

	res, err := h.Grpc.Shop.CreateShop(ctx, shop)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, errors.Wrap(err, "Failed to request GRPC shop").Error())
	}
	return res, nil
}

func (h *Handler) HandleEditShop(c *fiber.Ctx, shopId int32, editedShop *shopproto.Shop) (*shopproto.ShopResponse, error) {
	identity, ok := c.Locals(authutils.IdentityKey).(*authutils.Identity)
	if !ok {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "identity not found in context")
	}
	ctx, err := authutils.EmbedIdentityToContext(c.Context(), identity)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, errors.Wrap(err, "can't embed identity to grpc context").Error())
	}

	if shopId != editedShop.Id {
		return nil, fiber.NewError(fiber.StatusBadRequest, "shopId params is not the same as shopId in body")
	}

	res, err := h.Grpc.Shop.EditShop(ctx, &shopproto.EditShopRequest{
		EditedShop: editedShop,
	})
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, errors.Wrap(err, "Failed to request GRPC shop").Error())
	}
	return res, nil
}

func (h *Handler) HandleDeleteShop(c *fiber.Ctx, shopId int32) (*shopproto.DeleteResponse, error) {
	identity, ok := c.Locals(authutils.IdentityKey).(*authutils.Identity)
	if !ok {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "identity not found in context")
	}
	ctx, err := authutils.EmbedIdentityToContext(c.Context(), identity)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, errors.Wrap(err, "can't embed identity to grpc context").Error())
	}

	res, err := h.Grpc.Shop.DeleteShop(ctx, &shopproto.DeleteShopRequest{
		Id: shopId,
	})
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, errors.Wrap(err, "Failed to request GRPC shop").Error())
	}
	return res, nil
}

// ----- shop item -----

func (h *Handler) HandleGetShopItem(c *fiber.Ctx, shopId int32, shopItemId int32) (*shopproto.ShopItemResponse, error) {
	identity, ok := c.Locals(authutils.IdentityKey).(*authutils.Identity)
	if !ok {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "identity not found in context")
	}
	ctx, err := authutils.EmbedIdentityToContext(c.Context(), identity)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, errors.Wrap(err, "can't embed identity to grpc context").Error())
	}

	res, err := h.Grpc.Shop.GetShopItem(ctx, &shopproto.GetShopItemRequest{
		Id:     shopItemId,
		ShopId: shopId,
	})
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, errors.Wrap(err, "Failed to request GRPC shop").Error())
	}
	return res, nil
}

func (h *Handler) HandleGetAllShopItems(c *fiber.Ctx, shopId int32) (*shopproto.AllShopItemsResponse, error) {
	identity, ok := c.Locals(authutils.IdentityKey).(*authutils.Identity)
	if !ok {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "identity not found in context")
	}
	ctx, err := authutils.EmbedIdentityToContext(c.Context(), identity)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, errors.Wrap(err, "can't embed identity to grpc context").Error())
	}

	res, err := h.Grpc.Shop.GetAllShopItems(ctx, &shopproto.GetAllShopItemsRequest{
		ShopId: shopId,
	})
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, errors.Wrap(err, "Failed to request GRPC shop").Error())
	}
	return res, nil
}

func (h *Handler) HandleCreateShopItem(c *fiber.Ctx, shopItem *shopproto.CreateShopItemRequest) (*shopproto.ShopItemResponse, error) {
	identity, ok := c.Locals(authutils.IdentityKey).(*authutils.Identity)
	if !ok {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "identity not found in context")
	}
	ctx, err := authutils.EmbedIdentityToContext(c.Context(), identity)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, errors.Wrap(err, "can't embed identity to grpc context").Error())
	}

	res, err := h.Grpc.Shop.CreateShopItem(ctx, shopItem)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, errors.Wrap(err, "Failed to request GRPC shop").Error())
	}
	return res, nil
}

func (h *Handler) HandleEditShopItem(c *fiber.Ctx, shopId int32, shopItemId int32, editedItem *shopproto.ShopItem) (*shopproto.ShopItemResponse, error) {
	identity, ok := c.Locals(authutils.IdentityKey).(*authutils.Identity)
	if !ok {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "identity not found in context")
	}
	ctx, err := authutils.EmbedIdentityToContext(c.Context(), identity)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, errors.Wrap(err, "can't embed identity to grpc context").Error())
	}

	if shopItemId != editedItem.Id {
		return nil, fiber.NewError(fiber.StatusBadRequest, "shopItemId params is not the same as shopItemId in body")
	}

	res, err := h.Grpc.Shop.EditShopItem(ctx, &shopproto.EditShopItemRequest{
		EditedItem: editedItem,
	})
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, errors.Wrap(err, "Failed to request GRPC shop").Error())
	}
	return res, nil
}

func (h *Handler) HandleDeleteShopItem(c *fiber.Ctx, shopId int32, shopItemId int32) (*shopproto.DeleteResponse, error) {
	identity, ok := c.Locals(authutils.IdentityKey).(*authutils.Identity)
	if !ok {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "identity not found in context")
	}
	ctx, err := authutils.EmbedIdentityToContext(c.Context(), identity)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, errors.Wrap(err, "can't embed identity to grpc context").Error())
	}

	res, err := h.Grpc.Shop.DeleteShopItem(ctx, &shopproto.DeleteShopItemRequest{
		Id:     shopItemId,
		ShopId: shopId,
	})
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, errors.Wrap(err, "Failed to request GRPC shop").Error())
	}
	return res, nil
}
