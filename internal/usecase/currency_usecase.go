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

type CurrencyUseCase struct {
	DB                 *gorm.DB
	CurrencyRepository *repository.CurrencyRepository
	Validate           *validator.Validate
}

func NewCurrencyUseCase(DB *gorm.DB, currencyRepository *repository.CurrencyRepository, validate *validator.Validate) *CurrencyUseCase {
	return &CurrencyUseCase{DB: DB, CurrencyRepository: currencyRepository, Validate: validate}
}

func (u *CurrencyUseCase) Create(ctx *gin.Context, req *model.CreateCurrencyRequest) (*model.CurrencyResponse, *apperrors.AppError) {
	tx := u.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	var currency entity.Currency
	currency.Name = req.Name
	currency.Code = req.Code
	currency.ExchangeRate = req.ExchangeRate

	if err := u.CurrencyRepository.Create(tx, &currency); err != nil {
		return nil, apperrors.NewAppError(http.StatusInternalServerError, model.ErrorDetail{
			Field:     "general",
			ErrorCode: string(apperrors.ServerError),
			Param:     "",
			Message:   "can not create currency",
		})
	}

	if err := tx.Commit().Error; err != nil {
		return nil, apperrors.NewAppError(http.StatusInternalServerError, model.ErrorDetail{
			Field:     "general",
			ErrorCode: string(apperrors.CommitError),
			Param:     "",
			Message:   "can not create currency",
		})
	}

	return &model.CurrencyResponse{
		Currency: &currency,
	}, nil
}

func (u *CurrencyUseCase) Update(ctx *gin.Context, id string, req *model.UpdateCurrencyRequest) (*entity.Currency, *apperrors.AppError) {
	tx := u.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	var currency entity.Currency
	if err := u.CurrencyRepository.FindById(u.DB, &currency, id); err != nil {
		return nil, apperrors.NewAppError(http.StatusNotFound, model.ErrorDetail{
			Field:     "general",
			ErrorCode: string(apperrors.NotFound),
			Param:     "",
			Message:   "currency not found",
		})
	}

	currency.Name = req.Name
	currency.Code = req.Code
	currency.ExchangeRate = req.ExchangeRate

	if err := u.CurrencyRepository.Update(tx, &currency); err != nil {
		return nil, apperrors.NewAppError(http.StatusInternalServerError, model.ErrorDetail{
			Field:     "general",
			ErrorCode: string(apperrors.ServerError),
			Param:     "",
			Message:   "can not update currency",
		})
	}

	if err := tx.Commit().Error; err != nil {
		return nil, apperrors.NewAppError(http.StatusInternalServerError, model.ErrorDetail{
			Field:     "general",
			ErrorCode: string(apperrors.CommitError),
			Param:     "",
			Message:   "can not update currency",
		})
	}

	return &currency, nil
}

func (u *CurrencyUseCase) GetById(ctx *gin.Context, id string) (*entity.Currency, *apperrors.AppError) {
	var currency entity.Currency
	if err := u.CurrencyRepository.FindById(u.DB, &currency, id); err != nil {
		return nil, apperrors.NewAppError(http.StatusNotFound, model.ErrorDetail{
			Field:     "general",
			ErrorCode: string(apperrors.NotFound),
			Param:     "",
			Message:   "currency not found",
		})
	}
	return &currency, nil
}

func (u *CurrencyUseCase) GetAll(ctx *gin.Context) (*[]entity.Currency, *apperrors.AppError) {
	var currencies []entity.Currency
	if err := u.CurrencyRepository.FindAll(u.DB, &currencies); err != nil {
		return nil, apperrors.NewAppError(http.StatusInternalServerError, model.ErrorDetail{
			Field:     "general",
			ErrorCode: string(apperrors.ServerError),
			Param:     "",
			Message:   "can not get currencies",
		})
	}
	return &currencies, nil
}

func (u *CurrencyUseCase) Delete(ctx *gin.Context, id string) *apperrors.AppError {
	tx := u.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	var currency entity.Currency
	if err := u.CurrencyRepository.FindById(u.DB, &currency, id); err != nil {
		return apperrors.NewAppError(http.StatusNotFound, model.ErrorDetail{
			Field:     "general",
			ErrorCode: string(apperrors.NotFound),
			Param:     "",
			Message:   "currency not found",
		})
	}

	if err := u.CurrencyRepository.Delete(tx, &currency); err != nil {
		return apperrors.NewAppError(http.StatusInternalServerError, model.ErrorDetail{
			Field:     "general",
			ErrorCode: string(apperrors.ServerError),
			Param:     "",
			Message:   "can not delete currency",
		})
	}

	if err := tx.Commit().Error; err != nil {
		return apperrors.NewAppError(http.StatusInternalServerError, model.ErrorDetail{
			Field:     "general",
			ErrorCode: string(apperrors.CommitError),
			Param:     "",
			Message:   "can not delete currency",
		})
	}

	return nil
}
