package model

import (
	"time"
	"waizly/internal/entity"
)

type InvoiceResponse struct {
	ID         uint                  `json:"id"`
	Subject    string                `json:"subject,omitempty"`
	Status     string                `json:"status,omitempty"`
	IssueDate  time.Time             `json:"issue_date,omitempty"`
	DueDate    time.Time             `json:"due_date,omitempty"`
	User       entity.User           `json:"user,omitempty"`
	Tax        entity.Tax            `json:"tax,omitempty"`
	Customer   entity.Customer       `json:"customer,omitempty"`
	Currency   entity.Currency       `json:"currency,omitempty"`
	TotalItems int                   `json:"total_items,omitempty"`
	Items      []InvoiceItemResponse `json:"items,omitempty"`
}

type CreateInvoiceResponse struct {
	ID uint `json:"id"`
}

type UpdateInvoiceResponse struct {
	ID uint `json:"id"`
}

type CreateInvoiceRequest struct {
	Subject    string               `json:"subject" binding:"required"`
	Status     string               `json:"status" binding:"required,oneof=unpaid paid"`
	IssueDate  MyTime               `json:"issue_date" binding:"required" time_format:"2006-01-02"`
	DueDate    MyTime               `json:"due_date" binding:"required,gtefield=IssueDate" time_format:"2006-01-02"`
	CustomerId uint                 `json:"customer_id" binding:"required"`
	TaxId      uint                 `json:"tax_id" binding:"required"`
	CurrencyId uint                 `json:"currency_id" binding:"required"`
	Items      []InvoiceItemRequest `json:"items" binding:"required"`
}

type InvoiceItemRequest struct {
	ItemId    uint    `json:"item_id" binding:"required"`
	Quantity  float64 `json:"quantity" binding:"required,gt=0"`
	UnitPrice float64 `json:"unit_price" binding:"required,gt=0"`
}

type ListInvoiceRequest struct {
}

type SearchInvoiceRequest struct {
	ID         string `json:"id" binding:"required"`
	IssueDate  string `json:"issue_date" binding:"required" time_format:"2006-01-02"`
	DueDate    string `json:"due_date" binding:"required" time_format:"2006-01-02"`
	Subject    string `json:"subject" binding:"required"`
	TotalItems int    `json:"total_items" validate:"min=0"`
	Customer   string `json:"customer" binding:"required"`
	Status     string `json:"status" binding:"required,oneof=unpaid paid"`
	Page       int    `json:"page" validate:"min=1"`
	Size       int    `json:"size" validate:"min=1,max=100"`
}

type UpdateInvoiceRequest struct {
	Subject    string               `json:"subject" binding:"required"`
	Status     string               `json:"status" binding:"required,oneof=unpaid paid"`
	IssueDate  MyTime               `json:"issue_date" binding:"required" time_format:"2006-01-02"`
	DueDate    MyTime               `json:"due_date" binding:"required,gtefield=IssueDate" time_format:"2006-01-02"`
	CustomerId uint                 `json:"customer_id" binding:"required"`
	TaxId      uint                 `json:"tax_id" binding:"required"`
	CurrencyId uint                 `json:"currency_id" binding:"required"`
	Items      []InvoiceItemRequest `json:"items" binding:"required"`
}
