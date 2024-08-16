package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"waizly/internal/entity"
	"waizly/internal/helper"
	"waizly/internal/model"
	"waizly/internal/usecase"
)

type TaxController struct {
	TaxUseCase *usecase.TaxUseCase
}

func NewTaxController(taxUseCase *usecase.TaxUseCase) *TaxController {
	return &TaxController{TaxUseCase: taxUseCase}
}

func (c *TaxController) Index(ctx *gin.Context) {
	var errorOut []model.ErrorDetail

	res, err := c.TaxUseCase.GetAll(ctx)
	if err != nil {
		errorOut = append(errorOut, err.Details)
		helper.ErrorResponse[any, []model.ErrorDetail](ctx, err.Code, errorOut)
		return
	}

	helper.SuccessResponse[*[]entity.Tax, any](ctx, http.StatusOK, res)
}

func (c *TaxController) Create(ctx *gin.Context) {
	var errorOut []model.ErrorDetail
	var createTaxRequest model.CreateTaxRequest

	if err := ctx.ShouldBindBodyWithJSON(&createTaxRequest); err != nil {
		helper.ValidationErrorResponse(ctx, err)
		return
	}

	res, err := c.TaxUseCase.Create(ctx, &createTaxRequest)
	if err != nil {
		errorOut = append(errorOut, err.Details)
		helper.ErrorResponse[any, []model.ErrorDetail](ctx, err.Code, errorOut)
		return
	}

	helper.SuccessResponse[*entity.Tax, any](ctx, http.StatusOK, res)
}

func (c *TaxController) Show(ctx *gin.Context) {
	var errorOut []model.ErrorDetail
	id := ctx.Param("id")

	res, err := c.TaxUseCase.GetById(ctx, id)
	if err != nil {
		errorOut = append(errorOut, err.Details)
		helper.ErrorResponse[any, []model.ErrorDetail](ctx, err.Code, errorOut)
		return
	}

	helper.SuccessResponse[*entity.Tax, any](ctx, http.StatusOK, res)
}

func (c *TaxController) Update(ctx *gin.Context) {
	var errorOut []model.ErrorDetail
	var updateTaxRequest model.UpdateTaxRequest
	id := ctx.Param("id")

	if err := ctx.ShouldBindBodyWithJSON(&updateTaxRequest); err != nil {
		helper.ValidationErrorResponse(ctx, err)
		return
	}

	res, err := c.TaxUseCase.Update(ctx, id, &updateTaxRequest)
	if err != nil {
		errorOut = append(errorOut, err.Details)
		helper.ErrorResponse[any, []model.ErrorDetail](ctx, err.Code, errorOut)
		return
	}

	helper.SuccessResponse[*entity.Tax, any](ctx, http.StatusOK, res)
}

func (c *TaxController) Delete(ctx *gin.Context) {
	var errorOut []model.ErrorDetail
	id := ctx.Param("id")

	if err := c.TaxUseCase.Delete(ctx, id); err != nil {
		errorOut = append(errorOut, err.Details)
		helper.ErrorResponse[any, []model.ErrorDetail](ctx, err.Code, errorOut)
		return
	}

	helper.SuccessResponse[any, any](ctx, http.StatusNoContent, nil)
}
