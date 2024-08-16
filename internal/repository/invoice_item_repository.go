package repository

import (
	"gorm.io/gorm"
	"waizly/internal/entity"
)

type InvoiceItemRepository struct {
	Repository[entity.InvoiceItem]
}

func NewInvoiceItemRepository() *InvoiceItemRepository {
	return &InvoiceItemRepository{}
}

func (r *InvoiceItemRepository) FindByInvoiceId(db *gorm.DB, invoiceId any) ([]entity.InvoiceItem, error) {
	var invoiceItems []entity.InvoiceItem
	if err := db.Where("invoice_id = ?", invoiceId).Find(&invoiceItems).Error; err != nil {
		return nil, err
	}
	return invoiceItems, nil
}

func (r *InvoiceItemRepository) DeleteByInvoiceId(db *gorm.DB, invoiceId string) error {
	return db.Where("invoice_id = ?", invoiceId).Delete(&entity.InvoiceItem{}).Error
}
