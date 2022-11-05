package v1handler

import (
	"github.com/GrabItYourself/giys-backend/auth/pkg/authproto"
	"github.com/GrabItYourself/giys-backend/order/pkg/orderproto"
	"github.com/GrabItYourself/giys-backend/payment/pkg/paymentproto"
	"github.com/GrabItYourself/giys-backend/shop/pkg/shopproto"
	"github.com/GrabItYourself/giys-backend/user/pkg/userproto"
)

type Handler struct {
	Grpc *GrpcClients
}

type GrpcClients struct {
	User    userproto.UserServiceClient
	Auth    authproto.AuthClient
	Order   orderproto.OrderClient
	Shop    shopproto.ShopServiceClient
	Payment paymentproto.PaymentServiceClient
}

func NewHandler(
	grpc *GrpcClients,
) *Handler {
	return &Handler{
		Grpc: grpc,
	}
}
