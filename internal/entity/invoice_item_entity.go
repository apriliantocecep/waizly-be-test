package entity

import "time"

type InvoiceItem struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	InvoiceID uint `json:"invoice_id"`
	Invoice   Invoice
	ItemID    uint `json:"item_id"`
	Item      Item
	Quantity  float64   `json:"quantity" gorm:"type:decimal(10,2)"`
	Price     float64   `json:"price" gorm:"type:decimal(10,2)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (t *InvoiceItem) TableName() string {
	return "invoice_items"
}
