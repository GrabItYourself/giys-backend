package models

type Shop struct {
	Id          string  `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Name        string  `json:"name" gorm:"column:name;not null"`
	Image       *string `json:"image,omitempty" gorm:"column:image"`
	Description *string `json:"description,omitempty" gorm:"column:description"`
	Location    *string `json:"location,omitempty" gorm:"column:location"`
	Contact     *string `json:"contact,omitempty" gorm:"column:contact"`
}

func (Shop) TableName() string {
	return "shop"
}
