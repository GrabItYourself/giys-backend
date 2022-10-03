package models

type User struct {
	ID    string `json:"id" gorm:"column:id;primaryKey"`
	Role  string `json:"role" gorm:"column:role;type:role;not null"`
	Email string `json:"email" gorm:"column:email;not null"`
}
