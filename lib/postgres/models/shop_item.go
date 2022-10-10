package models

type ShopItem struct {
	Id     int32   `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	ShopId int32   `json:"shop_id" gorm:"column:shop_id;primaryKey;not null"`
	Name   string  `json:"name" gorm:"column:name;not null"`
	Image  *string `json:"image,omitempty" gorm:"column:image"`
	Price  int32   `json:"price" gorm:"column:price;not null"`
}

func (ShopItem) TableName() string {
	return "shop_item"
}
