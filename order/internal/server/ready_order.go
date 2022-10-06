package server

import (
	"context"

	"github.com/GrabItYourself/giys-backend/order/pkg/orderproto"
)

func (s *Server) ReadyOrder(ctx context.Context, in *orderproto.ReadyOrderRequest) (*orderproto.OrderResponse, error) {
	// var (
	// 	id    = in.Id
	// 	order = models.Order{}
	// )
	// if err := s.pg.Find(&order, models.Order{Id: int(id)}).Error; err != nil {
	// 	return nil, status.Error(postgres.InferCodeFromError(err), errors.Wrap(err, "can't find order").Error())
	// }
	// order.Status = "ready"
	// if err := s.pg.Save(&order).Error; err != nil {
	// 	return nil, status.Error(postgres.InferCodeFromError(err), errors.Wrap(err, "can't save order").Error())
	// }
	// return &orderproto.OrderResponse{
	// 	Id:     int32(order.Id),
	// 	UserId: order.UserId,
	// 	ShopId: order.ShopId,
	// 	Status: string(order.Status),
	// }, nil
	return nil, nil
}
