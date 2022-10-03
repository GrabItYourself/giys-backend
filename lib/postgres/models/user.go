package models

type RoleEnum string

const (
	AdminRole RoleEnum = "admin"
	UserRole  RoleEnum = "user"
)

type User struct {
	ID       string   `json:"id" gorm:"column:id;primaryKey"`
	Role     RoleEnum `json:"role" gorm:"column:role;type:role;not null"`
	Email    string   `json:"email" gorm:"column:email;not null"`
	GoogleID string   `json:"google_id" gorm:"column:google_id;not null"`
}

func (User) TableName() string {
	return "user"
}
