package models

type BankAccount struct {
	ShopId int32  `json:"shop_id" gorm:"column:shop_id;primaryKey"`
	Name   string `json:"name" gorm:"column:name"`
	Type   string `json:"type" gorm:"column:type"`
	Brand  string `json:"brand" gorm:"column:brand"`
	Number string `json:"number" gorm:"column:number"`
}

func (BankAccount) TableName() string {
	return "bank_account"
}
