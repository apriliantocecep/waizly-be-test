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

type ItemUseCase struct {
	DB             *gorm.DB
	ItemRepository *repository.ItemRepository
	Validate       *validator.Validate
}

func NewItemUseCase(DB *gorm.DB, itemRepository *repository.ItemRepository, validate *validator.Validate) *ItemUseCase {
	return &ItemUseCase{DB: DB, ItemRepository: itemRepository, Validate: validate}
}

func (u *ItemUseCase) Create(ctx *gin.Context, req *model.CreateItemRequest) (*entity.Item, *apperrors.AppError) {
	tx := u.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	var item entity.Item
	item.Name = req.Name
	item.Type = req.Type

	if err := u.ItemRepository.Create(tx, &item); err != nil {
		return nil, apperrors.NewAppError(http.StatusInternalServerError, model.ErrorDetail{
			Field:     "general",
			ErrorCode: string(apperrors.ServerError),
			Param:     "",
			Message:   "can not create item",
		})
	}

	if err := tx.Commit().Error; err != nil {
		return nil, apperrors.NewAppError(http.StatusInternalServerError, model.ErrorDetail{
			Field:     "general",
			ErrorCode: string(apperrors.CommitError),
			Param:     "",
			Message:   "can not create item",
		})
	}

	return &item, nil
}

func (u *ItemUseCase) GetAll(ctx *gin.Context) (*[]entity.Item, *apperrors.AppError) {
	var items []entity.Item
	if err := u.ItemRepository.FindAll(u.DB, &items); err != nil {
		return nil, apperrors.NewAppError(http.StatusInternalServerError, model.ErrorDetail{
			Field:     "general",
			ErrorCode: string(apperrors.ServerError),
			Param:     "",
			Message:   "can not get items",
		})
	}

	return &items, nil
}

func (u *ItemUseCase) GetById(ctx *gin.Context, id string) (*entity.Item, *apperrors.AppError) {
	var item entity.Item
	if err := u.ItemRepository.FindById(u.DB, &item, id); err != nil {
		return nil, apperrors.NewAppError(http.StatusNotFound, model.ErrorDetail{
			Field:     "general",
			ErrorCode: string(apperrors.NotFound),
			Param:     "",
			Message:   "item not found",
		})
	}

	return &item, nil
}

func (u *ItemUseCase) Update(ctx *gin.Context, id string, req *model.UpdateItemRequest) (*entity.Item, *apperrors.AppError) {
	tx := u.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	var item entity.Item
	if err := u.ItemRepository.FindById(u.DB, &item, id); err != nil {
		return nil, apperrors.NewAppError(http.StatusNotFound, model.ErrorDetail{
			Field:     "general",
			ErrorCode: string(apperrors.NotFound),
			Param:     "",
			Message:   "item not found",
		})
	}

	item.Name = req.Name
	item.Type = req.Type

	if err := u.ItemRepository.Update(tx, &item); err != nil {
		return nil, apperrors.NewAppError(http.StatusInternalServerError, model.ErrorDetail{
			Field:     "general",
			ErrorCode: string(apperrors.ServerError),
			Param:     "",
			Message:   "can not update item",
		})
	}

	if err := tx.Commit().Error; err != nil {
		return nil, apperrors.NewAppError(http.StatusInternalServerError, model.ErrorDetail{
			Field:     "general",
			ErrorCode: string(apperrors.CommitError),
			Param:     "",
			Message:   "can not update item",
		})
	}

	return &item, nil
}

func (u *ItemUseCase) Delete(ctx *gin.Context, id string) *apperrors.AppError {
	tx := u.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	var item entity.Item
	if err := u.ItemRepository.FindById(u.DB, &item, id); err != nil {
		return apperrors.NewAppError(http.StatusNotFound, model.ErrorDetail{
			Field:     "general",
			ErrorCode: string(apperrors.NotFound),
			Param:     "",
			Message:   "item not found",
		})
	}

	if err := u.ItemRepository.Delete(u.DB, &item); err != nil {
		return apperrors.NewAppError(http.StatusInternalServerError, model.ErrorDetail{
			Field:     "general",
			ErrorCode: string(apperrors.ServerError),
			Param:     "",
			Message:   "can not delete item",
		})
	}

	if err := tx.Commit().Error; err != nil {
		return apperrors.NewAppError(http.StatusInternalServerError, model.ErrorDetail{
			Field:     "general",
			ErrorCode: string(apperrors.CommitError),
			Param:     "",
			Message:   "can not delete item",
		})
	}

	return nil
}
