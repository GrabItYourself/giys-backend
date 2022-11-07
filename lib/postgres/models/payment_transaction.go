package models

import "time"

type PaymentTransaction struct {
	Id        int32     `json:"id" gorm:"column:id;primaryKey"`
	ShopId    int32     `json:"shop_id" gorm:"column:shop_id;not null"`
	OrderId   int32     `json:"order_id" gorm:"column:order_id;not null"`
	Amount    int       `json:"amount" gorm:"column:amount;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;not null"`
}

func (PaymentTransaction) TableName() string {
	return "payment_transaction"
}
