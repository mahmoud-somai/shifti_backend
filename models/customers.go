package models

type Customer struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	ForeignID   uint   `gorm:"not null" json:"foreign_id"`
	DateCreated string `gorm:"not null" json:"date_created"`
	Email       string `gorm:"not null" json:"email"`
	FirstName   string `gorm:"not null" json:"first_name"`
	LastName    string `gorm:"not null" json:"last_name"`
	Username    string `gorm:"not null" json:"username"`
	Password    string `gorm:"not null" json:"password"`
	AvatarURL   string `gorm:"not null" json:"avatar_url"`
	BillingId   uint   `gorm:"not null" json:"billing_id"`
}

func (Customer) TableName() string {
	return "customers"
}
