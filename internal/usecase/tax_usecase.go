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

type TaxUseCase struct {
	DB            *gorm.DB
	TaxRepository *repository.TaxRepository
	Validate      *validator.Validate
}

func NewTaxUseCase(DB *gorm.DB, taxRepository *repository.TaxRepository, validate *validator.Validate) *TaxUseCase {
	return &TaxUseCase{DB: DB, TaxRepository: taxRepository, Validate: validate}
}

func (u *TaxUseCase) Create(ctx *gin.Context, req *model.CreateTaxRequest) (*entity.Tax, *apperrors.AppError) {
	tx := u.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	var tax entity.Tax
	tax.Name = req.Name
	tax.Rate = req.Rate

	if err := u.TaxRepository.Create(tx, &tax); err != nil {
		return nil, apperrors.NewAppError(http.StatusInternalServerError, model.ErrorDetail{
			Field:     "general",
			ErrorCode: string(apperrors.ServerError),
			Param:     "",
			Message:   "can not create tax",
		})
	}

	if err := tx.Commit().Error; err != nil {
		return nil, apperrors.NewAppError(http.StatusInternalServerError, model.ErrorDetail{
			Field:     "general",
			ErrorCode: string(apperrors.CommitError),
			Param:     "",
			Message:   "can not create tax",
		})
	}

	return &tax, nil
}

func (u *TaxUseCase) GetAll(ctx *gin.Context) (*[]entity.Tax, *apperrors.AppError) {
	var taxes []entity.Tax
	if err := u.TaxRepository.FindAll(u.DB, &taxes); err != nil {
		return nil, apperrors.NewAppError(http.StatusInternalServerError, model.ErrorDetail{
			Field:     "general",
			ErrorCode: string(apperrors.ServerError),
			Param:     "",
			Message:   "can not get taxes",
		})
	}

	return &taxes, nil
}

func (u *TaxUseCase) GetById(ctx *gin.Context, id string) (*entity.Tax, *apperrors.AppError) {
	var tax entity.Tax
	if err := u.TaxRepository.FindById(u.DB, &tax, id); err != nil {
		return nil, apperrors.NewAppError(http.StatusNotFound, model.ErrorDetail{
			Field:     "general",
			ErrorCode: string(apperrors.NotFound),
			Param:     "",
			Message:   "tax not found",
		})
	}

	return &tax, nil
}

func (u *TaxUseCase) Update(ctx *gin.Context, id string, req *model.UpdateTaxRequest) (*entity.Tax, *apperrors.AppError) {
	tx := u.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	var tax entity.Tax
	if err := u.TaxRepository.FindById(u.DB, &tax, id); err != nil {
		return nil, apperrors.NewAppError(http.StatusNotFound, model.ErrorDetail{
			Field:     "general",
			ErrorCode: string(apperrors.NotFound),
			Param:     "",
			Message:   "tax not found",
		})
	}

	tax.Name = req.Name
	tax.Rate = req.Rate

	if err := u.TaxRepository.Update(tx, &tax); err != nil {
		return nil, apperrors.NewAppError(http.StatusInternalServerError, model.ErrorDetail{
			Field:     "general",
			ErrorCode: string(apperrors.ServerError),
			Param:     "",
			Message:   "can not update tax",
		})
	}

	if err := tx.Commit().Error; err != nil {
		return nil, apperrors.NewAppError(http.StatusInternalServerError, model.ErrorDetail{
			Field:     "general",
			ErrorCode: string(apperrors.CommitError),
			Param:     "",
			Message:   "can not update tax",
		})
	}

	return &tax, nil
}

func (u *TaxUseCase) Delete(ctx *gin.Context, id string) *apperrors.AppError {
	tx := u.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	var tax entity.Tax
	if err := u.TaxRepository.FindById(u.DB, &tax, id); err != nil {
		return apperrors.NewAppError(http.StatusNotFound, model.ErrorDetail{
			Field:     "general",
			ErrorCode: string(apperrors.NotFound),
			Param:     "",
			Message:   "tax not found",
		})
	}

	if err := u.TaxRepository.Delete(u.DB, &tax); err != nil {
		return apperrors.NewAppError(http.StatusInternalServerError, model.ErrorDetail{
			Field:     "general",
			ErrorCode: string(apperrors.ServerError),
			Param:     "",
			Message:   "can not delete tax",
		})
	}

	if err := tx.Commit().Error; err != nil {
		return apperrors.NewAppError(http.StatusInternalServerError, model.ErrorDetail{
			Field:     "general",
			ErrorCode: string(apperrors.CommitError),
			Param:     "",
			Message:   "can not delete tax",
		})
	}

	return nil
}
