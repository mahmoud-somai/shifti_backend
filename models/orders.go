package models

type Order struct {
	ForeignID          uint     `gorm:"not null" json:"foreign_id"`
	ID                 uint     `gorm:"primaryKey" json:"id"`
	ParentID           uint     `gorm:"not null" json:"parent_id"`
	OrderKey           string   `gorm:"not null" json:"order_key"`
	CreatedVia         string   `gorm:"not null" json:"created_via"`
	OrderVersion       string   `gorm:"not null" json:"order_version"`
	Status             string   `gorm:"not null" json:"status"`
	Currency           string   `gorm:"not null" json:"currency"`
	DateCreated        string   `gorm:"not null" json:"date_created"`
	DateModified       string   `gorm:"not null" json:"date_modified"`
	DiscountTotal      string   `gorm:"not null" json:"discount_total"`
	DiscountTax        string   `gorm:"not null" json:"discount_tax"`
	ShippingTotal      string   `gorm:"not null" json:"shipping_total"`
	ShippingTax        string   `gorm:"not null" json:"shipping_tax"`
	CartTax            string   `gorm:"not null" json:"cart_tax"`
	Total              string   `gorm:"not null" json:"total"`
	TotalTax           string   `gorm:"not null" json:"total_tax"`
	PricesIncludeTax   bool     `gorm:"not null" json:"prices_include_tax"`
	CustomerID         uint     `gorm:"not null" json:"customer_id"`
	CustomerIPAddress  string   `gorm:"not null" json:"customer_ip_address"`
	CustomerUserAgent  string   `gorm:"not null" json:"customer_user_agent"`
	CustomerNote       string   `gorm:"not null" json:"customer_note"`
	PaymentMethod      string   `gorm:"not null" json:"payment_method"`
	PaymentMethodTitle string   `gorm:"not null" json:"payment_method_title"`
	TransactionID      string   `gorm:"not null" json:"transaction_id"`
	DateCompleted      string   `gorm:"not null" json:"date_completed"`
	DatePaid           string   `gorm:"not null" json:"date_paid"`
	CartHash           string   `gorm:"not null" json:"cart_hash"`
	RefundLines        []string `gorm:"type:json" json:"refund_lines"`
	CouponLines        []string `gorm:"type:json" json:"coupon_lines"`
	FeeLines           []string `gorm:"type:json" json:"fee_lines"`

	ShippingLines struct {
		ShippingID          string `json:"shipping_id"`
		ShippingMethodTitle string `json:"shipping_method_title"`
		ShippingMethodID    string `json:"shipping_method_id"`
		ShippingMethodTotal string `json:"shipping_method_total"`
		ShippingMethodTax   string `json:"shipping_method_total_tax"`
		ShippingMethodTaxes string `json:"shipping_method_taxes"`
	} `gorm:"type:json" json:"shipping_lines"`

	TaxLines struct {
		TaxItemID               string `json:"tax_item_id"`
		TaxItemName             string `json:"tax_item_name"`
		TaxItemRateCode         string `json:"tax_item_rate_code"`
		TaxItemRateID           string `json:"tax_item_rate_id"`
		TaxItemRateLabel        string `json:"tax_item_rate_label"`
		TaxItemIsCompound       string `json:"tax_item_is_compound"`
		TaxItemCompound         string `json:"tax_item_compound"`
		TaxItemTaxTotal         string `json:"tax_item_tax_total"`
		TaxItemShippingTaxTotal string `json:"tax_item_shipping_tax_total"`
	} `gorm:"type:json" json:"tax_lines"`

	LineItems []string `gorm:"type:json" json:"line_items"`

	Billing struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Company   string `json:"company"`
		Address1  string `json:"address_1"`
		Address2  string `json:"address_2"`
		City      string `json:"city"`
		State     string `json:"state"`
		Postcode  string `json:"postcode"`
		Country   string `json:"country"`
		Email     string `json:"email"`
		Phone     string `json:"phone"`
	} `gorm:"type:json" json:"Billing"`

	Shipping struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Company   string `json:"company"`
		Address1  string `json:"address_1"`
		Address2  string `json:"address_2"`
		City      string `json:"city"`
		State     string `json:"state"`
		Postcode  string `json:"postcode"`
		Country   string `json:"country"`
		Phone     string `json:"phone"`
	} `gorm:"type:json" json:"Shipping"`
}

func (Order) TableName() string {
	return "orders"
}
