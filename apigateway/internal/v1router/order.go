package v1router

import (
	"strconv"

	"github.com/GrabItYourself/giys-backend/apigateway/internal/middlewares"
	"github.com/GrabItYourself/giys-backend/order/pkg/orderproto"
	"github.com/gofiber/fiber/v2"
)

func (r *Router) InitOrderRoute(basePath string) {
	f := r.Router.Group(basePath)

	// protect all paths below
	f.Use(middlewares.NewAccessTokenGuard(r.Handler.Grpc.Auth))

	f.Get("/:orderId", func(c *fiber.Ctx) error {
		shopIdStr, orderIdStr := c.Params("shopId"), c.Params("orderId")
		shopId, err := strconv.Atoi(shopIdStr)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "shopId is not a number")
		}
		orderId, err := strconv.Atoi(orderIdStr)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "orderId is not a number")
		}
		order, err := r.Handler.HandleGetOrder(c, int32(shopId), int32(orderId))
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		return c.JSON(order)
	})

	f.Post("/", func(c *fiber.Ctx) error {
		shopIdStr := c.Params("shopId")
		shopId, err := strconv.Atoi(shopIdStr)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "shopId is not a number")
		}
		reqBody := new(orderproto.CreateOrderRequest)
		if err := c.BodyParser(reqBody); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "items is not a valid json")
		}
		order, err := r.Handler.HandleCreateOrder(c, int32(shopId), reqBody.Items)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		return c.JSON(order)
	})

	f.Put("/:orderId", func(c *fiber.Ctx) error {
		shopIdStr, orderIdStr := c.Params("shopId"), c.Params("orderId")
		shopId, err := strconv.Atoi(shopIdStr)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "shopId is not a number")
		}
		orderId, err := strconv.Atoi(orderIdStr)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "orderId is not a number")
		}
		reqBody := new(orderproto.UpdateOrderRequest)
		if err := c.BodyParser(reqBody); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "items is not a valid json")
		}
		order, err := r.Handler.HandleUpdateOrder(c, int32(shopId), int32(orderId), reqBody.Items)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		return c.JSON(order)
	})

	f.Patch("/:orderId/complete", func(c *fiber.Ctx) error {
		shopIdStr, orderIdStr := c.Params("shopId"), c.Params("orderId")
		shopId, err := strconv.Atoi(shopIdStr)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "shopId is not a number")
		}
		orderId, err := strconv.Atoi(orderIdStr)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "orderId is not a number")
		}
		order, err := r.Handler.HandleCompleteOrder(c, int32(shopId), int32(orderId))
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		return c.JSON(order)
	})

	f.Patch("/:orderId/cancel", func(c *fiber.Ctx) error {
		shopIdStr, orderIdStr := c.Params("shopId"), c.Params("orderId")
		shopId, err := strconv.Atoi(shopIdStr)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "shopId is not a number")
		}
		orderId, err := strconv.Atoi(orderIdStr)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "orderId is not a number")
		}
		order, err := r.Handler.HandleCancelOrder(c, int32(shopId), int32(orderId))
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		return c.JSON(order)
	})

	f.Patch("/:orderId/ready", func(c *fiber.Ctx) error {
		shopIdStr, orderIdStr := c.Params("shopId"), c.Params("orderId")
		shopId, err := strconv.Atoi(shopIdStr)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "shopId is not a number")
		}
		orderId, err := strconv.Atoi(orderIdStr)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "orderId is not a number")
		}
		order, err := r.Handler.HandleReadyOrder(c, int32(orderId), int32(shopId))
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		return c.JSON(order)
	})

	f.Delete("/:orderId", func(c *fiber.Ctx) error {
		shopIdStr, orderIdStr := c.Params("shopId"), c.Params("orderId")
		shopId, err := strconv.Atoi(shopIdStr)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "shopId is not a number")
		}
		orderId, err := strconv.Atoi(orderIdStr)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "orderId is not a number")
		}
		order, err := r.Handler.HandleDeleteOrder(c, int32(shopId), int32(orderId))
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		return c.JSON(order)
	})
}
