package types

import (
	"github.com/GrabItYourself/giys-backend/payment/pkg/paymentproto"
	"github.com/GrabItYourself/giys-backend/shop/pkg/shopproto"
)

type CreateShopWithBankAccountRequest struct {
	*shopproto.CreateShopRequest
	BankAccount paymentproto.BankAccount `json:"bank_account"`
}
