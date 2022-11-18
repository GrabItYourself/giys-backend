package v1handler

import (
	"github.com/GrabItYourself/giys-backend/apigateway/internal/utils"
	"github.com/GrabItYourself/giys-backend/apigateway/internal/v1router/types"
	"github.com/GrabItYourself/giys-backend/auth/pkg/authutils"
	"github.com/GrabItYourself/giys-backend/payment/pkg/paymentproto"
	"github.com/GrabItYourself/giys-backend/shop/pkg/shopproto"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

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

func (h *Handler) HandleGetAllShops(c *fiber.Ctx) (*shopproto.AllShopsResponse, error) {
	identity, ok := c.Locals(authutils.IdentityKey).(*authutils.Identity)
	if !ok {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "identity not found in context")
	}
	ctx, err := authutils.EmbedIdentityToContext(c.Context(), identity)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, errors.Wrap(err, "can't embed identity to grpc context").Error())
	}

	res, err := h.Grpc.Shop.GetAllShops(ctx, &shopproto.GetAllShopsRequest{})
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, errors.Wrap(err, "Failed to request GRPC shop").Error())
	}
	return res, nil
}

func (h *Handler) HandleCreateShop(c *fiber.Ctx, shop *types.CreateShopWithBankAccountRequest) (*shopproto.ShopResponse, error) {
	identity, ok := c.Locals(authutils.IdentityKey).(*authutils.Identity)
	if !ok {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "identity not found in context")
	}
	ctx, err := authutils.EmbedIdentityToContext(c.Context(), identity)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, errors.Wrap(err, "can't embed identity to grpc context").Error())
	}

	if shop.Image != nil {
		image, err := utils.CompressImage(*shop.Image)
		if err != nil {
			return nil, fiber.NewError(fiber.StatusInternalServerError, errors.Wrap(err, "failed to compress image").Error())
		}
		shop.Image = &image
	}

	res, err := h.Grpc.Shop.CreateShop(ctx, shop.CreateShopRequest)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, errors.Wrap(err, "Failed to request GRPC shop").Error())
	}

	_, err = h.Grpc.Payment.RegisterRecipient(ctx, &paymentproto.RegisterRecipientRequest{
		ShopId:      res.Shop.Id,
		BankAccount: &shop.BankAccount,
	})
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, errors.Wrap(err, "Failed to request GRPC payment").Error())
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

	if editedShop.Image != nil {
		image, err := utils.CompressImage(*editedShop.Image)
		if err != nil {
			return nil, fiber.NewError(fiber.StatusInternalServerError, errors.Wrap(err, "failed to compress image").Error())
		}
		editedShop.Image = &image
	}

	res, err := h.Grpc.Shop.EditShop(ctx, &shopproto.EditShopRequest{
		EditedShop: editedShop,
	})
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, errors.Wrap(err, "Failed to request GRPC shop").Error())
	}
	return res, nil
}

func (h *Handler) HandleEditShopOwners(c *fiber.Ctx, shopId int32, ownerEmails []string) (*shopproto.ShopResponse, error) {
	identity, ok := c.Locals(authutils.IdentityKey).(*authutils.Identity)
	if !ok {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "identity not found in context")
	}
	ctx, err := authutils.EmbedIdentityToContext(c.Context(), identity)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, errors.Wrap(err, "can't embed identity to grpc context").Error())
	}

	res, err := h.Grpc.Shop.EditShopOwners(ctx, &shopproto.EditShopOwnersRequest{
		ShopId:      shopId,
		OwnerEmails: ownerEmails,
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
