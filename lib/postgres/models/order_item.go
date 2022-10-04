package models

type OrderItem struct {
	Id         string `json:"id" gorm:"column:id;primaryKey"`
	ShopId     string `json:"shop_id" gorm:"column:shop_id;not null"`
	ShopItemId string `json:"shop_item_id" gorm:"column:shop_item_id;not null"`
	Quantity   int    `json:"quantity" gorm:"column:quantity;not null"`
	Note       string `json:"note" gorm:"column:note"`
}

func (OrderItem) TableName() string {
	return "order_item"
}
