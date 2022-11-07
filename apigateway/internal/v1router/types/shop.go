package types

type AddBankAccountRequest struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Number string `json:"number"`
	Brand  string `json:"brand"`
}
