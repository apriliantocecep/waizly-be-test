package repository

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"waizly/internal/entity"
	"waizly/internal/model"
)

type InvoiceRepository struct {
	Repository[entity.Invoice]
}

func NewInvoiceRepository() *InvoiceRepository {
	return &InvoiceRepository{}
}

func (r *InvoiceRepository) FindByIdWithPreloads(db *gorm.DB, invoice *entity.Invoice, id string) error {
	return db.Where("id = ?", id).Preload(clause.Associations).Find(invoice).Error
}

func (r *InvoiceRepository) Search(db *gorm.DB, req *model.SearchInvoiceRequest) ([]entity.Invoice, int64, error) {
	var invoices []entity.Invoice
	if err := db.Preload(clause.Associations).Scopes(r.Filter(req)).Offset((req.Page - 1) * req.Size).Limit(req.Size).Find(&invoices).Error; err != nil {
		return nil, 0, err
	}

	var total int64
	if err := db.Model(&entity.Invoice{}).Scopes(r.Filter(req)).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return invoices, total, nil
}

func (r *InvoiceRepository) Filter(req *model.SearchInvoiceRequest) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if id := req.ID; id != "" {
			tx = tx.Where("id = ?", id)
		}
		if subject := req.Subject; subject != "" {
			tx = tx.Where("subject LIKE ?", "%"+subject+"%")
		}
		if customer := req.Customer; customer != "" {
			tx = tx.Joins("JOIN customers ON invoices.customer_id = customers.id").
				Where("customers.name LIKE ?", "%"+customer+"%")
		}
		if status := req.Status; status != "" {
			tx = tx.Where("status = ?", status)
		}
		if issueDate := req.IssueDate; issueDate != "" {
			tx = tx.Where("issue_date = ?", issueDate)
		}
		if dueDate := req.DueDate; dueDate != "" {
			tx = tx.Where("due_date = ?", dueDate)
		}
		if totalItems := req.TotalItems; totalItems != 0 {
			tx = tx.Joins("LEFT JOIN invoice_items ON invoices.id = invoice_items.invoice_id").
				Group("invoices.id").
				Having("COUNT(invoice_items.id) = ?", totalItems)
		}
		return tx
	}
}

func (r *InvoiceRepository) DeleteWithAssociations(db *gorm.DB, invoice *entity.Invoice) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("invoice_id = ?", invoice.ID).Delete(&entity.InvoiceItem{}).Error; err != nil {
			return err
		}
		if err := tx.Delete(invoice).Error; err != nil {
			return err
		}
		return nil
	})
}

func (r *InvoiceRepository) DeleteItems(db *gorm.DB, invoiceId string) error {
	return db.Where("invoice_id = ?", invoiceId).Delete(&entity.InvoiceItem{}).Error
}
