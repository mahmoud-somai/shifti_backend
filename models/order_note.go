package models

import "time"

type OrderNote struct {
	ID        uint      `gorm:"primaryKey" json:"id_order_note"`
	ForeignId uint      `gorm:"not null" json:"id_foreign"`
	Date      time.Time `gorm:"not null" json:"note_date"`
	Author    string    `gorm:"size:255;not null" json:"author"`
	Content   string    `gorm:"type:text;not null" json:"content"`
}

func (OrderNote) TableName() string {
	return "order_notes"
}
