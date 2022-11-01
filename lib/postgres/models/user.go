package models

type RoleEnum string

func (r RoleEnum) MarshalBinary() ([]byte, error) {
	return []byte(r), nil
}

const (
	AdminRole RoleEnum = "ADMIN"
	UserRole  RoleEnum = "USER"
)

type User struct {
	Id                     string   `json:"id" gorm:"column:id;primaryKey;default:gen_random_uuid()"`
	Role                   RoleEnum `json:"role" gorm:"column:role;type:role;not null;default:USER"`
	Email                  string   `json:"email" gorm:"column:email;not null"`
	GoogleId               string   `json:"google_id" gorm:"column:google_id;not null"`
	OmiseCustomerId        *string  `json:"omise_customer_id" gorm:"column:omise_customer_id"`
	DefaultPaymentMethodId *string  `json:"default_payment_method_id" gorm:"column:default_payment_method_id"`
}

func (User) TableName() string {
	return "user"
}
