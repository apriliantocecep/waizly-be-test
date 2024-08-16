package converter

import (
	"waizly/internal/entity"
	"waizly/internal/model"
)

func InvoiceToResponse(invoice *entity.Invoice) *model.InvoiceResponse {
	invoiceItems := make([]model.InvoiceItemResponse, len(invoice.InvoiceItems))

	for i, item := range invoice.InvoiceItems {
		invoiceItems[i] = *InvoiceItemToResponse(&item)
	}
	return &model.InvoiceResponse{
		ID:         invoice.ID,
		Subject:    invoice.Subject,
		Status:     invoice.Status,
		IssueDate:  invoice.IssueDate,
		DueDate:    invoice.DueDate,
		User:       invoice.User,
		Tax:        invoice.Tax,
		Customer:   invoice.Customer,
		Currency:   invoice.Currency,
		TotalItems: len(invoice.InvoiceItems),
		Items:      invoiceItems,
	}
}
