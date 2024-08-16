package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"log"
	"net/http"
	"sync"
	"time"
	"waizly/internal/entity"
	"waizly/internal/model"
	"waizly/internal/model/converter"
	"waizly/internal/repository"
	"waizly/pkg/apperrors"
)

type InvoiceUseCase struct {
	DB                    *gorm.DB
	InvoiceRepository     *repository.InvoiceRepository
	InvoiceItemRepository *repository.InvoiceItemRepository
	CustomerRepository    *repository.CustomerRepository
	TaxRepository         *repository.TaxRepository
	CurrencyRepository    *repository.CurrencyRepository
	UserRepository        *repository.UserRepository
	Validate              *validator.Validate
}

func NewInvoiceUseCase(DB *gorm.DB, invoiceRepository *repository.InvoiceRepository, invoiceItemRepository *repository.InvoiceItemRepository, customerRepository *repository.CustomerRepository, taxRepository *repository.TaxRepository, currencyRepository *repository.CurrencyRepository, userRepository *repository.UserRepository, validate *validator.Validate) *InvoiceUseCase {
	return &InvoiceUseCase{DB: DB, InvoiceRepository: invoiceRepository, InvoiceItemRepository: invoiceItemRepository, CustomerRepository: customerRepository, TaxRepository: taxRepository, CurrencyRepository: currencyRepository, UserRepository: userRepository, Validate: validate}
}

func (u *InvoiceUseCase) Create(ctx *gin.Context, req *model.CreateInvoiceRequest) (*model.CreateInvoiceResponse, *apperrors.AppError) {
	tx := u.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	userId, _ := ctx.Get("userId")
	var invoice entity.Invoice
	var customer entity.Customer
	var tax entity.Tax
	var currency entity.Currency

	// find customer
	if err := u.CustomerRepository.FindById(tx, &customer, req.CustomerId); err != nil {
		return nil, apperrors.NewAppError(http.StatusNotFound, model.ErrorDetail{
			Field:     "customer_id",
			ErrorCode: string(apperrors.NotFound),
			Param:     "",
			Message:   "customer not found",
		})
	}

	// find tax
	if err := u.TaxRepository.FindById(tx, &tax, req.TaxId); err != nil {
		return nil, apperrors.NewAppError(http.StatusNotFound, model.ErrorDetail{
			Field:     "tax_id",
			ErrorCode: string(apperrors.NotFound),
			Param:     "",
			Message:   "tax not found",
		})
	}

	// find currency
	if err := u.CurrencyRepository.FindById(tx, &currency, req.CurrencyId); err != nil {
		return nil, apperrors.NewAppError(http.StatusNotFound, model.ErrorDetail{
			Field:     "currency_id",
			ErrorCode: string(apperrors.NotFound),
			Param:     "",
			Message:   "currency not found",
		})
	}

	invoice.Subject = req.Subject
	invoice.Status = req.Status
	invoice.IssueDate = time.Time(req.IssueDate)
	invoice.DueDate = time.Time(req.DueDate)
	invoice.CustomerId = customer.ID
	invoice.TaxId = tax.ID
	invoice.CurrencyId = currency.ID
	invoice.UserId = userId.(uint)

	if err := u.InvoiceRepository.Create(tx, &invoice); err != nil {
		return nil, apperrors.NewAppError(http.StatusInternalServerError, model.ErrorDetail{
			Field:     "general",
			ErrorCode: string(apperrors.ServerError),
			Param:     "",
			Message:   "can not create invoice",
		})
	}

	// save invoice items
	var wg sync.WaitGroup

	for _, item := range req.Items {
		wg.Add(1)

		go func(item model.InvoiceItemRequest) {
			defer wg.Done()

			var invoiceItem entity.InvoiceItem
			invoiceItem.InvoiceID = invoice.ID
			invoiceItem.ItemID = item.ItemId
			invoiceItem.Quantity = item.Quantity
			invoiceItem.Price = item.UnitPrice
			if err := u.InvoiceItemRepository.Create(tx, &invoiceItem); err != nil {
				log.Printf("Error inserting invoice item: %v", err)
				return
			}
		}(item)
	}

	wg.Wait()

	if err := tx.Commit().Error; err != nil {
		return nil, apperrors.NewAppError(http.StatusInternalServerError, model.ErrorDetail{
			Field:     "general",
			ErrorCode: string(apperrors.CommitError),
			Param:     "",
			Message:   "can not create invoice",
		})
	}

	return &model.CreateInvoiceResponse{ID: invoice.ID}, nil
}

func (u *InvoiceUseCase) GetById(ctx *gin.Context, id string) (*model.InvoiceResponse, *apperrors.AppError) {
	var invoice entity.Invoice
	if err := u.InvoiceRepository.FindByIdWithPreloads(u.DB, &invoice, id); err != nil {
		return nil, apperrors.NewAppError(http.StatusNotFound, model.ErrorDetail{
			Field:     "general",
			ErrorCode: string(apperrors.NotFound),
			Param:     "",
			Message:   "invoice not found",
		})
	}

	return converter.InvoiceToResponse(&invoice), nil
}

func (u *InvoiceUseCase) Search(ctx *gin.Context, req *model.SearchInvoiceRequest) ([]model.InvoiceResponse, int64, *apperrors.AppError) {
	invoices, total, err := u.InvoiceRepository.Search(u.DB, req)
	if err != nil {
		return nil, 0, apperrors.NewAppError(http.StatusInternalServerError, model.ErrorDetail{
			Field:     "general",
			ErrorCode: string(apperrors.ServerError),
			Param:     "",
			Message:   "can not search invoices",
		})
	}

	responses := make([]model.InvoiceResponse, len(invoices))
	for i, invoice := range invoices {
		responses[i] = *converter.InvoiceToResponse(&invoice)
	}

	return responses, total, nil
}

func (u *InvoiceUseCase) Delete(ctx *gin.Context, id string) *apperrors.AppError {
	tx := u.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	var invoice entity.Invoice
	if err := u.InvoiceRepository.FindById(tx, &invoice, id); err != nil {
		return apperrors.NewAppError(http.StatusNotFound, model.ErrorDetail{
			Field:     "general",
			ErrorCode: string(apperrors.NotFound),
			Param:     "",
			Message:   "invoice not found",
		})
	}

	if err := u.InvoiceRepository.DeleteWithAssociations(tx, &invoice); err != nil {
		return apperrors.NewAppError(http.StatusInternalServerError, model.ErrorDetail{
			Field:     "general",
			ErrorCode: string(apperrors.ServerError),
			Param:     "",
			Message:   "can not delete invoice with items",
		})
	}

	if err := tx.Commit().Error; err != nil {
		return apperrors.NewAppError(http.StatusInternalServerError, model.ErrorDetail{
			Field:     "general",
			ErrorCode: string(apperrors.CommitError),
			Param:     "",
			Message:   "can not delete invoice",
		})
	}

	return nil
}

func (u *InvoiceUseCase) Update(ctx *gin.Context, id string, req *model.UpdateInvoiceRequest) (*model.UpdateInvoiceResponse, *apperrors.AppError) {
	tx := u.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	userId, _ := ctx.Get("userId")
	var invoice entity.Invoice
	var customer entity.Customer
	var tax entity.Tax
	var currency entity.Currency

	// find invoice
	if err := u.InvoiceRepository.FindById(tx, &invoice, id); err != nil {
		return nil, apperrors.NewAppError(http.StatusNotFound, model.ErrorDetail{
			Field:     "general",
			ErrorCode: string(apperrors.NotFound),
			Param:     "",
			Message:   "invoice not found",
		})
	}

	// find customer
	if err := u.CustomerRepository.FindById(tx, &customer, req.CustomerId); err != nil {
		return nil, apperrors.NewAppError(http.StatusNotFound, model.ErrorDetail{
			Field:     "customer_id",
			ErrorCode: string(apperrors.NotFound),
			Param:     "",
			Message:   "customer not found",
		})
	}

	// find tax
	if err := u.TaxRepository.FindById(tx, &tax, req.TaxId); err != nil {
		return nil, apperrors.NewAppError(http.StatusNotFound, model.ErrorDetail{
			Field:     "tax_id",
			ErrorCode: string(apperrors.NotFound),
			Param:     "",
			Message:   "tax not found",
		})
	}

	// find currency
	if err := u.CurrencyRepository.FindById(tx, &currency, req.CurrencyId); err != nil {
		return nil, apperrors.NewAppError(http.StatusNotFound, model.ErrorDetail{
			Field:     "currency_id",
			ErrorCode: string(apperrors.NotFound),
			Param:     "",
			Message:   "currency not found",
		})
	}

	// update invoice
	invoice.Subject = req.Subject
	invoice.Status = req.Status
	invoice.IssueDate = time.Time(req.IssueDate)
	invoice.DueDate = time.Time(req.DueDate)
	invoice.CustomerId = customer.ID
	invoice.TaxId = tax.ID
	invoice.CurrencyId = currency.ID
	invoice.UserId = userId.(uint)

	if err := u.InvoiceRepository.Update(tx, &invoice); err != nil {
		return nil, apperrors.NewAppError(http.StatusInternalServerError, model.ErrorDetail{
			Field:     "general",
			ErrorCode: string(apperrors.ServerError),
			Param:     "",
			Message:   "can not update invoice",
		})
	}

	// delete existing invoice items
	errDelete := u.InvoiceItemRepository.DeleteByInvoiceId(tx, id)
	if errDelete != nil {
		return nil, apperrors.NewAppError(http.StatusInternalServerError, model.ErrorDetail{
			Field:     "general",
			ErrorCode: string(apperrors.ServerError),
			Param:     "",
			Message:   "can not delete invoice items",
		})
	} else {
		// insert new items
		var wg sync.WaitGroup

		for _, item := range req.Items {
			wg.Add(1)

			go func(item model.InvoiceItemRequest) {
				defer wg.Done()

				var invoiceItem entity.InvoiceItem
				invoiceItem.InvoiceID = invoice.ID
				invoiceItem.ItemID = item.ItemId
				invoiceItem.Quantity = item.Quantity
				invoiceItem.Price = item.UnitPrice
				if err := u.InvoiceItemRepository.Create(tx, &invoiceItem); err != nil {
					log.Printf("Error updating invoice item: %v", err)
					return
				}
			}(item)
		}

		wg.Wait()
	}

	if err := tx.Commit().Error; err != nil {
		return nil, apperrors.NewAppError(http.StatusInternalServerError, model.ErrorDetail{
			Field:     "general",
			ErrorCode: string(apperrors.CommitError),
			Param:     "",
			Message:   "can not update invoice",
		})
	}

	return &model.UpdateInvoiceResponse{
		ID: invoice.ID,
	}, nil
}
