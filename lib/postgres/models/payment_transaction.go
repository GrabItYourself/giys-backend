package models

import "time"

type PaymentTransaction struct {
	Id        uint      `json:"id" gorm:"column:id;primaryKey"`
	ShopId    string    `json:"shop_id" gorm:"column:shop_id;not null"`
	OrderId   string    `json:"order_id" gorm:"column:order_id;not null"`
	Amount    int       `json:"amount" gorm:"column:amount;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;not null"`
}

func (PaymentTransaction) TableName() string {
	return "payment_transaction"
}
