package entity

import "time"

type Invoice struct {
	ID           uint          `json:"id" gorm:"primaryKey"`
	Subject      string        `json:"subject"`
	Status       string        `json:"status"`
	IssueDate    time.Time     `json:"issue_date"`
	DueDate      time.Time     `json:"due_date"`
	CustomerId   uint          `json:"customer_id"`
	Customer     Customer      `json:"customer" gorm:"foreignKey:customer_id;references:id"`
	TaxId        uint          `json:"tax_id"`
	Tax          Tax           `json:"tax" gorm:"foreignKey:tax_id;references:id"`
	CurrencyId   uint          `json:"currency_id"`
	Currency     Currency      `json:"currency" gorm:"foreignKey:currency_id;references:id"`
	UserId       uint          `json:"user_id" gorm:"column:user_id"`
	User         User          `json:"user" gorm:"foreignKey:user_id;references:id"`
	InvoiceItems []InvoiceItem `json:"invoice_items"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
}

func (t *Invoice) TableName() string {
	return "invoices"
}
