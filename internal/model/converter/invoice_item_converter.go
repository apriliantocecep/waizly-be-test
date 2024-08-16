package converter

import (
	"waizly/internal/entity"
	"waizly/internal/model"
)

func InvoiceItemToResponse(invoiceItem *entity.InvoiceItem) *model.InvoiceItemResponse {
	return &model.InvoiceItemResponse{
		ID:       invoiceItem.ID,
		Quantity: invoiceItem.Quantity,
		Price:    invoiceItem.Price,
		ItemId:   invoiceItem.ItemID,
	}
}
