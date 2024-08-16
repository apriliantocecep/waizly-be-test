package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"waizly/internal/entity"
	"waizly/internal/helper"
	"waizly/internal/model"
	"waizly/internal/usecase"
)

type CurrencyController struct {
	CurrencyUseCase *usecase.CurrencyUseCase
}

func NewCurrencyController(currencyUseCase *usecase.CurrencyUseCase) *CurrencyController {
	return &CurrencyController{CurrencyUseCase: currencyUseCase}
}

func (c *CurrencyController) Index(ctx *gin.Context) {
	var errorOut []model.ErrorDetail
	res, err := c.CurrencyUseCase.GetAll(ctx)
	if err != nil {
		errorOut = append(errorOut, err.Details)
		helper.ErrorResponse[any, []model.ErrorDetail](ctx, err.Code, errorOut)
		return
	}
	helper.SuccessResponse[*[]entity.Currency, any](ctx, http.StatusOK, res)
}

func (c *CurrencyController) Show(ctx *gin.Context) {
	var errorOut []model.ErrorDetail
	id := ctx.Param("id")
	res, err := c.CurrencyUseCase.GetById(ctx, id)
	if err != nil {
		errorOut = append(errorOut, err.Details)
		helper.ErrorResponse[any, []model.ErrorDetail](ctx, err.Code, errorOut)
		return
	}
	helper.SuccessResponse[*entity.Currency, any](ctx, http.StatusOK, res)
}

func (c *CurrencyController) Create(ctx *gin.Context) {
	var errorOut []model.ErrorDetail
	var createCurrencyRequest model.CreateCurrencyRequest
	if err := ctx.ShouldBindBodyWithJSON(&createCurrencyRequest); err != nil {
		helper.ValidationErrorResponse(ctx, err)
		return
	}
	res, err := c.CurrencyUseCase.Create(ctx, &createCurrencyRequest)
	if err != nil {
		errorOut = append(errorOut, err.Details)
		helper.ErrorResponse[any, []model.ErrorDetail](ctx, err.Code, errorOut)
		return
	}
	helper.SuccessResponse[*model.CurrencyResponse, any](ctx, http.StatusOK, res)
}

func (c *CurrencyController) Update(ctx *gin.Context) {
	var errorOut []model.ErrorDetail
	var updateCurrencyRequest model.UpdateCurrencyRequest
	id := ctx.Param("id")
	if err := ctx.ShouldBindBodyWithJSON(&updateCurrencyRequest); err != nil {
		helper.ValidationErrorResponse(ctx, err)
		return
	}
	res, err := c.CurrencyUseCase.Update(ctx, id, &updateCurrencyRequest)
	if err != nil {
		errorOut = append(errorOut, err.Details)
		helper.ErrorResponse[any, []model.ErrorDetail](ctx, err.Code, errorOut)
		return
	}
	helper.SuccessResponse[*entity.Currency, any](ctx, http.StatusOK, res)
}

func (c *CurrencyController) Delete(ctx *gin.Context) {
	var errorOut []model.ErrorDetail
	id := ctx.Param("id")
	if err := c.CurrencyUseCase.Delete(ctx, id); err != nil {
		errorOut = append(errorOut, err.Details)
		helper.ErrorResponse[any, []model.ErrorDetail](ctx, err.Code, errorOut)
		return
	}
	helper.SuccessResponse[any, any](ctx, http.StatusNoContent, nil)
}
