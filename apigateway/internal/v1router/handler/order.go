package v1handler

import (
	"github.com/GrabItYourself/giys-backend/auth/pkg/authutils"
	"github.com/GrabItYourself/giys-backend/order/pkg/orderproto"
	"github.com/GrabItYourself/giys-backend/payment/pkg/paymentproto"
	"github.com/GrabItYourself/giys-backend/user/pkg/userproto"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

func (h *Handler) HandleGetOrder(c *fiber.Ctx, shopId, orderId int32) (*orderproto.OrderResponse, error) {
	identity, ok := c.Locals(authutils.IdentityKey).(*authutils.Identity)
	if !ok {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "identity not found in context")
	}
	ctx, err := authutils.EmbedIdentityToContext(c.Context(), identity)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, errors.Wrap(err, "can't embed identity to grpc context").Error())
	}

	res, err := h.Grpc.Order.GetOrder(ctx, &orderproto.GetOrderRequest{
		OrderId: orderId,
		ShopId:  shopId,
	})
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, errors.Wrap(err, "Failed to request GRPC order").Error())
	}
	return res, nil
}

func (h *Handler) HandleCreateOrder(c *fiber.Ctx, shopId int32, items []*orderproto.OrderItem) (*orderproto.OrderResponse, error) {
	identity, ok := c.Locals(authutils.IdentityKey).(*authutils.Identity)
	if !ok {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "identity not found in context")
	}
	ctx, err := authutils.EmbedIdentityToContext(c.Context(), identity)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, errors.Wrap(err, "can't embed identity to grpc context").Error())
	}

	me, err := h.Grpc.User.Me(ctx, &userproto.MeReq{})
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, errors.Wrap(err, "Failed to request GRPC user").Error())
	}

	orderRes, err := h.Grpc.Order.CreateOrder(ctx, &orderproto.CreateOrderRequest{
		ShopId: shopId,
		UserId: me.User.Id,
		Items:  items,
	})
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, errors.Wrap(err, "Failed to request GRPC order").Error())
	}
	return orderRes, nil
}

func (h *Handler) HandleUpdateOrder(c *fiber.Ctx, shopId int32, orderId int32, items []*orderproto.OrderItem) (*orderproto.OrderResponse, error) {
	identity, ok := c.Locals(authutils.IdentityKey).(*authutils.Identity)
	if !ok {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "identity not found in context")
	}
	ctx, err := authutils.EmbedIdentityToContext(c.Context(), identity)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, errors.Wrap(err, "can't embed identity to grpc context").Error())
	}

	me, err := h.Grpc.User.Me(ctx, &userproto.MeReq{})
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, errors.Wrap(err, "Failed to request GRPC user").Error())
	}

	orderRes, err := h.Grpc.Order.UpdateOrder(ctx, &orderproto.UpdateOrderRequest{
		OrderId: orderId,
		ShopId:  shopId,
		UserId:  me.User.Id,
		Items:   items,
	})
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, errors.Wrap(err, "Failed to request GRPC order").Error())
	}
	return orderRes, nil
}

func (h *Handler) HandleDeleteOrder(c *fiber.Ctx, shopId int32, orderId int32) (*orderproto.OrderResponse, error) {
	identity, ok := c.Locals(authutils.IdentityKey).(*authutils.Identity)
	if !ok {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "identity not found in context")
	}
	ctx, err := authutils.EmbedIdentityToContext(c.Context(), identity)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, errors.Wrap(err, "can't embed identity to grpc context").Error())
	}

	me, err := h.Grpc.User.Me(ctx, &userproto.MeReq{})
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, errors.Wrap(err, "Failed to request GRPC user").Error())
	}

	orderRes, err := h.Grpc.Order.DeleteOrder(ctx, &orderproto.DeleteOrderRequest{
		OrderId: orderId,
		ShopId:  shopId,
		UserId:  me.User.Id,
	})
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, errors.Wrap(err, "Failed to request GRPC order").Error())
	}
	return orderRes, nil
}

func (h *Handler) HandleCompleteOrder(c *fiber.Ctx, shopId int32, orderId int32) (*orderproto.OrderResponse, error) {
	identity, ok := c.Locals(authutils.IdentityKey).(*authutils.Identity)
	if !ok {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "identity not found in context")
	}
	ctx, err := authutils.EmbedIdentityToContext(c.Context(), identity)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, errors.Wrap(err, "can't embed identity to grpc context").Error())
	}

	_, err = h.Grpc.Payment.Pay(ctx, &paymentproto.PayRequest{
		ShopId:  shopId,
		OrderId: orderId,
	})
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, errors.Wrap(err, "Failed to request GRPC Pay").Error())
	}

	orderRes, err := h.Grpc.Order.CompleteOrder(ctx, &orderproto.CompleteOrderRequest{
		OrderId: orderId,
		ShopId:  shopId,
	})
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, errors.Wrap(err, "Failed to request GRPC order").Error())
	}
	return orderRes, nil
}

func (h *Handler) HandleReadyOrder(c *fiber.Ctx, shopId int32, orderId int32) (*orderproto.OrderResponse, error) {
	identity, ok := c.Locals(authutils.IdentityKey).(*authutils.Identity)
	if !ok {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "identity not found in context")
	}
	ctx, err := authutils.EmbedIdentityToContext(c.Context(), identity)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, errors.Wrap(err, "can't embed identity to grpc context").Error())
	}

	orderRes, err := h.Grpc.Order.ReadyOrder(ctx, &orderproto.ReadyOrderRequest{
		OrderId: orderId,
		ShopId:  shopId,
	})
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, errors.Wrap(err, "Failed to request GRPC order").Error())
	}
	return orderRes, nil
}

func (h *Handler) HandleCancelOrder(c *fiber.Ctx, shopId int32, orderId int32) (*orderproto.OrderResponse, error) {
	identity, ok := c.Locals(authutils.IdentityKey).(*authutils.Identity)
	if !ok {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "identity not found in context")
	}
	ctx, err := authutils.EmbedIdentityToContext(c.Context(), identity)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, errors.Wrap(err, "can't embed identity to grpc context").Error())
	}

	orderRes, err := h.Grpc.Order.CancelOrder(ctx, &orderproto.CancelOrderRequest{
		OrderId: orderId,
		ShopId:  shopId,
	})
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, errors.Wrap(err, "Failed to request GRPC order").Error())
	}
	return orderRes, nil
}
