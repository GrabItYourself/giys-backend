package models

type OrderItem struct {
	Id int `json:"id" gorm:"column:id;primaryKey"`

	OrderId int   `json:"order_id" gorm:"column:order_id;not null"`
	Order   Order `json:"order" gorm:"foreignKey:OrderId"`

	ShopId int  `json:"shop_id" gorm:"column:shop_id;not null"`
	Shop   Shop `json:"shop" gorm:"foreignKey:ShopId"`

	ShopItemId int      `json:"shop_item_id" gorm:"column:shop_item_id;not null"`
	ShopItem   ShopItem `json:"shop_item" gorm:"foreignKey:ShopItemId"`

	Quantity int    `json:"quantity" gorm:"column:quantity;not null"`
	Note     string `json:"note" gorm:"column:note"`
}

func (OrderItem) TableName() string {
	return "order_item"
}
