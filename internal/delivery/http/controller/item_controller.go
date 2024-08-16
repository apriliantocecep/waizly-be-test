package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"waizly/internal/entity"
	"waizly/internal/helper"
	"waizly/internal/model"
	"waizly/internal/usecase"
)

type ItemController struct {
	ItemUseCase *usecase.ItemUseCase
}

func NewItemController(itemUseCase *usecase.ItemUseCase) *ItemController {
	return &ItemController{ItemUseCase: itemUseCase}
}

func (c *ItemController) Index(ctx *gin.Context) {
	var errorOut []model.ErrorDetail

	res, err := c.ItemUseCase.GetAll(ctx)
	if err != nil {
		errorOut = append(errorOut, err.Details)
		helper.ErrorResponse[any, []model.ErrorDetail](ctx, err.Code, errorOut)
		return
	}

	helper.SuccessResponse[*[]entity.Item, any](ctx, http.StatusOK, res)
}

func (c *ItemController) Create(ctx *gin.Context) {
	var errorOut []model.ErrorDetail
	var createItemRequest model.CreateItemRequest

	if err := ctx.ShouldBindBodyWithJSON(&createItemRequest); err != nil {
		helper.ValidationErrorResponse(ctx, err)
		return
	}

	res, err := c.ItemUseCase.Create(ctx, &createItemRequest)
	if err != nil {
		errorOut = append(errorOut, err.Details)
		helper.ErrorResponse[any, []model.ErrorDetail](ctx, err.Code, errorOut)
		return
	}

	helper.SuccessResponse[*entity.Item, any](ctx, http.StatusOK, res)
}

func (c *ItemController) Show(ctx *gin.Context) {
	var errorOut []model.ErrorDetail
	id := ctx.Param("id")

	res, err := c.ItemUseCase.GetById(ctx, id)
	if err != nil {
		errorOut = append(errorOut, err.Details)
		helper.ErrorResponse[any, []model.ErrorDetail](ctx, err.Code, errorOut)
		return
	}

	helper.SuccessResponse[*entity.Item, any](ctx, http.StatusOK, res)
}

func (c *ItemController) Update(ctx *gin.Context) {
	var errorOut []model.ErrorDetail
	var updateItemRequest model.UpdateItemRequest
	id := ctx.Param("id")

	if err := ctx.ShouldBindBodyWithJSON(&updateItemRequest); err != nil {
		helper.ValidationErrorResponse(ctx, err)
		return
	}

	res, err := c.ItemUseCase.Update(ctx, id, &updateItemRequest)
	if err != nil {
		errorOut = append(errorOut, err.Details)
		helper.ErrorResponse[any, []model.ErrorDetail](ctx, err.Code, errorOut)
		return
	}

	helper.SuccessResponse[*entity.Item, any](ctx, http.StatusOK, res)
}

func (c *ItemController) Delete(ctx *gin.Context) {
	var errorOut []model.ErrorDetail
	id := ctx.Param("id")

	if err := c.ItemUseCase.Delete(ctx, id); err != nil {
		errorOut = append(errorOut, err.Details)
		helper.ErrorResponse[any, []model.ErrorDetail](ctx, err.Code, errorOut)
		return
	}

	helper.SuccessResponse[any, any](ctx, http.StatusNoContent, nil)
}
