package v1handler

import (
	"github.com/GrabItYourself/giys-backend/apigateway/internal/utils"
	"github.com/GrabItYourself/giys-backend/auth/pkg/authutils"
	"github.com/GrabItYourself/giys-backend/shop/pkg/shopproto"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

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

	if shopItem.Image != nil {
		image, err := utils.CompressImage(*shopItem.Image)
		if err != nil {
			return nil, fiber.NewError(fiber.StatusInternalServerError, errors.Wrap(err, "failed to compress image").Error())
		}
		shopItem.Image = &image
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

	if editedItem.Image != nil {
		image, err := utils.CompressImage(*editedItem.Image)
		if err != nil {
			return nil, fiber.NewError(fiber.StatusInternalServerError, errors.Wrap(err, "failed to compress image").Error())
		}
		editedItem.Image = &image
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
