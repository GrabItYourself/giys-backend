package models

type Shop struct {
	ID          string `json:"id" gorm:"column:id;primaryKey"`
	Name        string `json:"name" gorm:"column:name;not null"`
	Image       string `json:"image" gorm:"column:image"`
	Description string `json:"description" gorm:"column:description"`
	Location    string `json:"location" gorm:"column:location"`
	Contact     string `json:"contact" gorm:"column:contact"`
}
