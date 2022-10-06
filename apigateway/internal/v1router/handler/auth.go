package v1handler

import (
	"context"

	"github.com/GrabItYourself/giys-backend/auth/pkg/authproto"
)

func (h *Handler) HandleGoogleOAuthCallback(ctx context.Context, code string) (*authproto.ExchangeAuthCodeResp, error) {
	resp, err := h.Grpc.Auth.ExchangeAuthCode(ctx, &authproto.ExchangeAuthCodeReq{
		AuthCode: code,
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
