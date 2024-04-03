package models

type Taxes struct {
	ID        uint     `gorm:"primaryKey" json:"id"`
	ForeignId uint     `gorm:"not null" json:"foreign_id"`
	Country   string   `gorm:"size:2;not null" json:"country"`
	State     string   `gorm:"size:255;not null" json:"state"`
	Postcodes []string `gorm:"type:json" json:"postcodes"`
	Cities    []string `gorm:"type:json" json:"cities"`
	Rate      string   `gorm:"not null" json:"rate"`
	Name      string   `gorm:"not null" json:"name"`
	Priority  string   `gorm:"not null" json:"priority"`
	Compound  bool     `gorm:"not null" json:"compound"`
	Shipping  bool     `gorm:"not null" json:"shipping"`
	Order     int      `gorm:"not null" json:"order"`
	Class     string   `gorm:"not null" json:"class"`
}

func (Taxes) TableName() string {
	return "taxes"
}
