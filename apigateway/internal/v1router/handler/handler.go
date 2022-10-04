package v1handler

import "github.com/GrabItYourself/giys-backend/user/pkg/libproto"

type Handler struct {
	UserClient libproto.UserServiceClient
}

func NewHandler(
	userClient libproto.UserServiceClient,
) *Handler {
	return &Handler{
		UserClient: userClient,
	}
}
