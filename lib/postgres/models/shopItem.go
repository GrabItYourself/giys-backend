package models

type ShopItem struct {
	Id     string  `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	ShopID string  `json:"shop_id" gorm:"column:shop_id;not null"`
	Name   string  `json:"name" gorm:"column:name;not null"`
	Image  *string `json:"image,omitempty" gorm:"column:image"`
	Price  string  `json:"price" gorm:"column:price;not null"`
}

func (ShopItem) TableName() string {
	return "shop_item"
}
