package models

type ShopItem struct {
	ID     string `json:"id" gorm:"column:id;primaryKey"`
	ShopID string `json:"shop_id" gorm:"column:shop_id;not null"`
	Name   string `json:"name" gorm:"column:name;not null"`
	Image  string `json:"image" gorm:"column:image"`
	Price  string `json:"price" gorm:"column:price;not null"`
}

func (ShopItem) TableName() string {
	return "shop_item"
}
