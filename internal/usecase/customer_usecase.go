package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"net/http"
	"waizly/internal/entity"
	"waizly/internal/model"
	"waizly/internal/repository"
	"waizly/pkg/apperrors"
)

type CustomerUseCase struct {
	DB                 *gorm.DB
	CustomerRepository *repository.CustomerRepository
	Validate           *validator.Validate
}

func NewCustomerUseCase(DB *gorm.DB, customerRepository *repository.CustomerRepository, validate *validator.Validate) *CustomerUseCase {
	return &CustomerUseCase{DB: DB, CustomerRepository: customerRepository, Validate: validate}
}

func (u *CustomerUseCase) Create(ctx *gin.Context, req *model.CreateCustomerRequest) (*entity.Customer, *apperrors.AppError) {
	tx := u.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	var customer entity.Customer
	customer.Name = req.Name
	customer.Status = req.Status
	customer.Address = req.Address

	if err := u.CustomerRepository.Create(tx, &customer); err != nil {
		return nil, apperrors.NewAppError(http.StatusInternalServerError, model.ErrorDetail{
			Field:     "general",
			ErrorCode: string(apperrors.ServerError),
			Param:     "",
			Message:   "can not create customer",
		})
	}

	if err := tx.Commit().Error; err != nil {
		return nil, apperrors.NewAppError(http.StatusInternalServerError, model.ErrorDetail{
			Field:     "general",
			ErrorCode: string(apperrors.CommitError),
			Param:     "",
			Message:   "can not create customer",
		})
	}

	return &customer, nil
}

func (u *CustomerUseCase) GetAll(ctx *gin.Context) (*[]entity.Customer, *apperrors.AppError) {
	var customers []entity.Customer
	if err := u.CustomerRepository.FindAll(u.DB, &customers); err != nil {
		return nil, apperrors.NewAppError(http.StatusInternalServerError, model.ErrorDetail{
			Field:     "general",
			ErrorCode: string(apperrors.ServerError),
			Param:     "",
			Message:   "can not get customers",
		})
	}

	return &customers, nil
}

func (u *CustomerUseCase) GetById(ctx *gin.Context, id string) (*entity.Customer, *apperrors.AppError) {
	var customer entity.Customer
	if err := u.CustomerRepository.FindById(u.DB, &customer, id); err != nil {
		return nil, apperrors.NewAppError(http.StatusNotFound, model.ErrorDetail{
			Field:     "general",
			ErrorCode: string(apperrors.NotFound),
			Param:     "",
			Message:   "customer not found",
		})
	}

	return &customer, nil
}

func (u *CustomerUseCase) Update(ctx *gin.Context, id string, req *model.UpdateCustomerRequest) (*entity.Customer, *apperrors.AppError) {
	tx := u.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	var customer entity.Customer
	if err := u.CustomerRepository.FindById(u.DB, &customer, id); err != nil {
		return nil, apperrors.NewAppError(http.StatusNotFound, model.ErrorDetail{
			Field:     "general",
			ErrorCode: string(apperrors.NotFound),
			Param:     "",
			Message:   "customer not found",
		})
	}

	customer.Name = req.Name
	customer.Status = req.Status
	customer.Address = req.Address

	if err := u.CustomerRepository.Update(tx, &customer); err != nil {
		return nil, apperrors.NewAppError(http.StatusInternalServerError, model.ErrorDetail{
			Field:     "general",
			ErrorCode: string(apperrors.ServerError),
			Param:     "",
			Message:   "can not update customer",
		})
	}

	if err := tx.Commit().Error; err != nil {
		return nil, apperrors.NewAppError(http.StatusInternalServerError, model.ErrorDetail{
			Field:     "general",
			ErrorCode: string(apperrors.CommitError),
			Param:     "",
			Message:   "can not update customer",
		})
	}

	return &customer, nil
}

func (u *CustomerUseCase) Delete(ctx *gin.Context, id string) *apperrors.AppError {
	tx := u.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	var customer entity.Customer
	if err := u.CustomerRepository.FindById(u.DB, &customer, id); err != nil {
		return apperrors.NewAppError(http.StatusNotFound, model.ErrorDetail{
			Field:     "general",
			ErrorCode: string(apperrors.NotFound),
			Param:     "",
			Message:   "customer not found",
		})
	}

	if err := u.CustomerRepository.Delete(u.DB, &customer); err != nil {
		return apperrors.NewAppError(http.StatusInternalServerError, model.ErrorDetail{
			Field:     "general",
			ErrorCode: string(apperrors.ServerError),
			Param:     "",
			Message:   "can not delete customer",
		})
	}

	if err := tx.Commit().Error; err != nil {
		return apperrors.NewAppError(http.StatusInternalServerError, model.ErrorDetail{
			Field:     "general",
			ErrorCode: string(apperrors.CommitError),
			Param:     "",
			Message:   "can not delete customer",
		})
	}

	return nil
}
