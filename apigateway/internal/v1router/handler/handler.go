package v1handler

import (
	"github.com/GrabItYourself/giys-backend/user/pkg/userproto"
)

type Handler struct {
	UserClient userproto.UserServiceClient
}

func NewHandler(
	userClient userproto.UserServiceClient,
) *Handler {
	return &Handler{
		UserClient: userClient,
	}
}
