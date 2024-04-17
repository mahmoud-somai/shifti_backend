package models

type OrderNote struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	ForeignId string `gorm:"not null" json:"note_id"`
	Date      string `gorm:"size:255;not null" json:"note_date"`
	Author    string `gorm:"size:255;not null" json:"note_author"`
	Content   string `gorm:"not null" json:"note_content"`
}

func (OrderNote) TableName() string {
	return "order_notes"
}
