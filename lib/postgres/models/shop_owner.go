package models

type ShopOwner struct {
	ShopId string `json:"shop_id" gorm:"column:shop_id;not null"`
	Shop   Shop   `json:"shop" gorm:"foreignKey:ShopId;references:Id"`
	UserId string `json:"user_id" gorm:"column:user_id;not null"`
	User   User   `json:"user" gorm:"foreignKey:UserId;references:Id"`
}

func (ShopOwner) TableName() string {
	return "shop_owner"
}
