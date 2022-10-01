package models

type role string

const (
	AdminRole role = "admin"
	UserRole  role = "user"
)

type User struct {
	ID    string `gorm:"column:id;primaryKey"`
	Role  role   `gorm:"column:role;type:role;not null"`
	Email string `gorm:"column:email;not null"`
}
