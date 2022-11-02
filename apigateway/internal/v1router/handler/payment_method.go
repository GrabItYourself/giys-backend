package v1handler

import (
	"github.com/GrabItYourself/giys-backend/auth/pkg/authutils"
	"github.com/GrabItYourself/giys-backend/payment/pkg/paymentproto"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

func (h *Handler) HandleGetMyPaymentMethods(c *fiber.Ctx) ([]*paymentproto.PaymentMethod, error) {
	identity, ok := c.Locals(authutils.IdentityKey).(*authutils.Identity)
	if !ok {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "identity not found in context")
	}
	ctx, err := authutils.EmbedIdentityToContext(c.Context(), identity)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, errors.Wrap(err, "can't embed identity to grpc context").Error())
	}

	res, err := h.Grpc.Payment.GetMyPaymentMethods(ctx, &paymentproto.GetMyPaymentMethodsRequest{})
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, errors.Wrap(err, "Failed to request GRPC GetMyPaymentMethods").Error())
	}
	return res.PaymentMethods, nil
}

func (h *Handler) HandleAddPaymentMethod(c *fiber.Ctx, name string, cardNumber string, expirationMonth int32, expirationyear int32) (*paymentproto.PaymentMethod, error) {
	identity, ok := c.Locals(authutils.IdentityKey).(*authutils.Identity)
	if !ok {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "identity not found in context")
	}
	ctx, err := authutils.EmbedIdentityToContext(c.Context(), identity)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, errors.Wrap(err, "can't embed identity to grpc context").Error())
	}

	res, err := h.Grpc.Payment.AuthorizeCard(ctx, &paymentproto.AuthorizeCardRequest{
		Name:            name,
		CardNumber:      cardNumber,
		ExpirationMonth: expirationMonth,
		ExpirationYear:  expirationyear,
	})
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, errors.Wrap(err, "Failed to request GRPC AuthorizeCard").Error())
	}
	return res.PaymentMethod, nil
}

func (h *Handler) HandleSetDefaultPaymentMethod(c *fiber.Ctx, paymentMethodId int32) (*paymentproto.UpdateDefaultPaymentMethodResponse, error) {
	identity, ok := c.Locals(authutils.IdentityKey).(*authutils.Identity)
	if !ok {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "identity not found in context")
	}
	ctx, err := authutils.EmbedIdentityToContext(c.Context(), identity)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, errors.Wrap(err, "can't embed identity to grpc context").Error())
	}

	res, err := h.Grpc.Payment.UpdateDefaultPaymentMethod(ctx, &paymentproto.UpdateDefaultPaymentMethodRequest{
		PaymentMethodId: paymentMethodId,
	})
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, errors.Wrap(err, "Failed to request GRPC UpdateDefaultPaymentMethod").Error())
	}
	return res, nil
}
