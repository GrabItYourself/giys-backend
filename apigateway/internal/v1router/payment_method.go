package v1router

import (
	"strconv"

	"github.com/GrabItYourself/giys-backend/apigateway/internal/middlewares"
	"github.com/GrabItYourself/giys-backend/payment/pkg/paymentproto"
	"github.com/gofiber/fiber/v2"
)

func (r *Router) InitPaymentMethodRoutes(basePath string) {
	f := r.Router.Group(basePath)

	// protect all paths below
	f.Use(middlewares.NewAccessTokenGuard(r.Handler.Grpc.Auth))

	f.Get("/", func(c *fiber.Ctx) error {
		paymentMethods, err := r.Handler.HandleGetAllPaymentMethods(c)
		if err != nil {
			return err
		}
		return c.JSON(paymentMethods)
	})

	f.Post("/", func(c *fiber.Ctx) error {
		reqBody := new(paymentproto.AuthorizeCardRequest)
		if err := c.BodyParser(reqBody); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "items is not a valid json")
		}

		paymentMethod, err := r.Handler.HandleAddPaymentMethod(c, reqBody.Name, reqBody.CardNumber, reqBody.ExpirationMonth, reqBody.ExpirationYear)
		if err != nil {
			return err
		}
		return c.JSON(paymentMethod)
	})

	f.Patch("/:paymentMethodId/setDefault", func(c *fiber.Ctx) error {
		paymentMethodIdStr := c.Params("paymentMethodId")

		paymentMethodId, err := strconv.Atoi(paymentMethodIdStr)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "paymentMethodId is not a number")
		}

		res, err := r.Handler.HandleSetDefaultPaymentMethod(c, int32(paymentMethodId))
		if err != nil {
			return err
		}
		return c.JSON(res)
	})
}
