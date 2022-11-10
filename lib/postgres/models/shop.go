package models

type Shop struct {
	Id               int32   `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Name             string  `json:"name" gorm:"column:name;not null;unique"`
	Image            *string `json:"image,omitempty" gorm:"column:image"`
	Description      *string `json:"description,omitempty" gorm:"column:description"`
	Location         *string `json:"location,omitempty" gorm:"column:location"`
	Contact          *string `json:"contact,omitempty" gorm:"column:contact"`
	OmiseResipientId *string `json:"omise_recipient_id" gorm:"column:omise_recipient_id"`
	OwnerId          string  `json:"owner_id" gorm:"column:owner_id;not null"`
	Owner            User    `json:"owner" gorm:"foreignKey:OwnerId;references:Id"`
}

func (Shop) TableName() string {
	return "shop"
}
