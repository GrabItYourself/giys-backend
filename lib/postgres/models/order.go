package models

type Order struct {
	Id     int    `json:"id" gorm:"column:id;primaryKey"`
	UserId string `json:"user_id" gorm:"column:user_id;type:role;not null"`
	ShopId string `json:"shop_id" gorm:"column:shop_id;not null"`
	Status string `json:"status" gorm:"column:status;not null"`
}

func (Order) TableName() string {
	return "order"
}
