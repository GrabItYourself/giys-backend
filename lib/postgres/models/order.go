package models

type OrderStatus string

const (
	InQueueStatus   OrderStatus = "IN_QUEUE"
	ReadyStatus     OrderStatus = "READY"
	CompletedStatus OrderStatus = "COMPLETED"
	CanceledStatus  OrderStatus = "CANCELED"
)

type Order struct {
	Id     int32       `json:"id" gorm:"column:id;primaryKey"`
	UserId string      `json:"user_id" gorm:"column:user_id;type:role;not null"`
	ShopId int32       `json:"shop_id" gorm:"column:shop_id;not null"`
	Status OrderStatus `json:"status" gorm:"column:status;not null;default:IN_QUEUE"`

	Items []OrderItem `json:"items" gorm:"foreignKey:OrderId"`
}

func (Order) TableName() string {
	return "order"
}
