package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type StringSlice []string
type IntSlice []int

func (ss StringSlice) Value() (driver.Value, error) {
	return json.Marshal(ss)
}

func (is IntSlice) Value() (driver.Value, error) {
	return json.Marshal(is)
}

func (is *IntSlice) Scan(value interface{}) error {
	if value == nil {
		*is = nil
		return nil
	}
	var data []byte
	switch val := value.(type) {
	case string:
		data = []byte(val)
	case []byte:
		data = val
	default:
		return fmt.Errorf("unsupported type for IntSlice: %T", value)
	}
	return json.Unmarshal(data, is)
}

func (ss *StringSlice) Scan(value interface{}) error {
	if value == nil {
		*ss = nil
		return nil
	}
	var data []byte
	switch val := value.(type) {
	case string:
		data = []byte(val)
	case []byte:
		data = val
	default:
		return fmt.Errorf("unsupported type for StringSlice: %T", value)
	}
	return json.Unmarshal(data, ss)
}

type Product struct {
	ID                uint        `gorm:"primaryKey" json:"product_id"`
	ForeignID         uint        `gorm:"not null" json:"id"`
	Name              string      `gorm:"not null" json:"name"`
	Slug              string      `gorm:"not null" json:"slug"`
	Permalink         string      `gorm:"not null" json:"permalink"`
	DateCreated       string      `gorm:"not null" json:"date_created"`
	DateModified      string      `gorm:"not null" json:"date_modified"`
	Type              string      `gorm:"not null" json:"type"`
	Status            string      `gorm:"not null" json:"status"`
	CatalogVisibility string      `gorm:"not null" json:"catalog_visibility"`
	Description       string      `gorm:"not null" json:"description"`
	ShortDescription  string      `gorm:"not null" json:"short_description"`
	SKU               string      `gorm:"not null" json:"sku"`
	Price             string      `gorm:"not null" json:"price"`
	RegularPrice      string      `gorm:"not null" json:"regular_price"`
	SalePrice         string      `gorm:"not null" json:"sale_price"`
	TaxStatus         string      `gorm:"not null" json:"tax_status"`
	TaxClass          string      `gorm:"not null" json:"tax_class"`
	ManageStock       bool        `gorm:"not null" json:"manage_stock"`
	StockQuantity     int         `gorm:"not null" json:"stock_quantity"`
	StockStatus       string      `gorm:"not null" json:"stock_status"`
	Backorders        string      `gorm:"not null" json:"backorders"`
	BackordersAllowed bool        `gorm:"not null" json:"backorders_allowed"`
	SoldIndividually  bool        `gorm:"not null" json:"sold_individually"`
	Weight            string      `gorm:"not null" json:"weight"`
	ShippingRequired  bool        `gorm:"not null" json:"shipping_required"`
	ShippingTaxable   bool        `gorm:"not null" json:"shipping_taxable"`
	ShippingClass     string      `gorm:"not null" json:"shipping_class"`
	ShippingClassID   int         `gorm:"not null" json:"shipping_class_id"`
	ReviewsAllowed    bool        `gorm:"not null" json:"reviews_allowed"`
	AverageRating     string      `gorm:"not null" json:"average_rating"`
	RatingCount       int         `gorm:"not null" json:"rating_count"`
	RelatedIDs        StringSlice `gorm:"type:json" json:"related_ids"`
	UpsellIDs         IntSlice    `gorm:"type:json" json:"upsell_ids"`
	CrossSellIDs      StringSlice `gorm:"type:json" json:"cross_sell_ids"`
	DefaultAttributes StringSlice `gorm:"type:json" json:"default_attributes"`
	Variations        IntSlice    `gorm:"type:json" json:"variations"`
	GroupedProducts   IntSlice    `gorm:"type:json" json:"grouped_products"`
	ParentID          int         `gorm:"not null" json:"parent_id"`
	PurchaseNote      string      `gorm:"not null" json:"purchase_note"`
	Categories        []Cat       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"categories"`
	Tags              []Tag       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"tags"`
	Images            []Image     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"images"`
	Attributes        []Attribute `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"attributes"`
	MenuOrder         int         `gorm:"not null" json:"menu_order"`
	Dimensions        []Dimension `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"dimensions"`

	DateOnSaleFrom string        `json:"date_on_sale_from"`
	DateOnSaleTo   string        `json:"date_on_sale_to"`
	OnSale         bool          `gorm:"not null" json:"on_sale"`
	Purchasable    bool          `gorm:"not null" json:"purchasable"`
	TotalSales     int           `gorm:"not null" json:"total_sales"`
	Virtual        bool          `gorm:"not null" json:"virtual"`
	Downloadable   bool          `gorm:"not null" json:"downloadable"`
	Downloads      []StringSlice `gorm:"type:json" json:"downloads"`
	DownloadLimit  int           `gorm:"not null" json:"download_limit"`
	DownloadExpiry int           `gorm:"not null" json:"download_expiry"`
}

type Cat struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	ForeignID uint   `gorm:"not null" json:"id_category"`
	Name      string `gorm:"not null" json:"name_category"`
	Slug      string `gorm:"not null" json:"slug_category"`
	ProductID uint   `gorm:"not null;index" json:"product_id"`
}

type Tag struct {
	ID        uint   `gorm:"primaryKey" json:"tag_id"`
	ForeignID uint   `gorm:"not null" json:"id"`
	Name      string `gorm:"not null" json:"name"`
	Slug      string `gorm:"not null" json:"slug"`
	ProductID uint   `gorm:"not null;index" json:"product_id"`
}

type Image struct {
	ID           uint   `gorm:"primaryKey" json:"id_image"`
	ForeignID    uint   `gorm:"not null" json:"id"`
	Name         string `gorm:"not null" json:"name"`
	Alt          string `gorm:"not null" json:"alt"`
	Src          string `gorm:"not null" json:"src"`
	DateCreated  string `json:"date_created"`
	DateModified string `json:"date_modified"`
	ProductID    uint   `gorm:"not null;index" json:"product_id"`
}

type Attribute struct {
	ID        uint   `gorm:"primaryKey" json:"id_attribute"`
	ForeignID uint   `gorm:"not null" json:"id"`
	Name      string `gorm:"not null" json:"name"`
	Position  int    `gorm:"not null" json:"position"`
	ProductID uint   `gorm:"not null;index" json:"product_id"`
}

type Dimension struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Height    string `json:"height"`
	Width     string `json:"width"`
	Length    string `json:"length"`
	ProductID uint   `gorm:"not null;index" json:"product_id"`
}
