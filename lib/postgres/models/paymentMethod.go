package models

type PaymentMethod struct {
	Id          uint   `json:"id" gorm:"column:id;primaryKey"`
	UserId      string `json:"user_id" gorm:"column:user_id;not null"`
	OmiseCardId string `json:"omise_card_id" gorm:"column:omise_card_id;not null"`
}

func (PaymentMethod) TableName() string {
	return "payment_method"
}
