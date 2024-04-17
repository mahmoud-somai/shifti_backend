package models

type Category struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	ForeignId   uint   `gorm:"not null" json:"foreign_id"`
	Name        string `gorm:"not null" json:"name"`
	Slug        string `gorm:"not null" json:"slug"`
	Parent      uint   `gorm:"not null" json:"parent"`
	Description string `gorm:"not null" json:"description"`
	Display     string `gorm:"not null" json:"display"`
	MenuOrder   int    `gorm:"not null" json:"menu_order"`
	Count       int    `gorm:"not null" json:"count"`
}

func (Category) TableName() string {
	return "categories"
}
