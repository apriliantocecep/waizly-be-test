package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"waizly/internal/entity"
	"waizly/internal/helper"
	"waizly/internal/model"
	"waizly/internal/usecase"
)

type CustomerController struct {
	CustomerUseCase *usecase.CustomerUseCase
}

func NewCustomerController(customerUseCase *usecase.CustomerUseCase) *CustomerController {
	return &CustomerController{CustomerUseCase: customerUseCase}
}

func (c *CustomerController) Index(ctx *gin.Context) {
	var errorOut []model.ErrorDetail

	res, err := c.CustomerUseCase.GetAll(ctx)
	if err != nil {
		errorOut = append(errorOut, err.Details)
		helper.ErrorResponse[any, []model.ErrorDetail](ctx, err.Code, errorOut)
		return
	}

	helper.SuccessResponse[*[]entity.Customer, any](ctx, http.StatusOK, res)
}

func (c *CustomerController) Create(ctx *gin.Context) {
	var errorOut []model.ErrorDetail
	var createCustomerRequest model.CreateCustomerRequest

	if err := ctx.ShouldBindBodyWithJSON(&createCustomerRequest); err != nil {
		helper.ValidationErrorResponse(ctx, err)
		return
	}

	res, err := c.CustomerUseCase.Create(ctx, &createCustomerRequest)
	if err != nil {
		errorOut = append(errorOut, err.Details)
		helper.ErrorResponse[any, []model.ErrorDetail](ctx, err.Code, errorOut)
		return
	}

	helper.SuccessResponse[*entity.Customer, any](ctx, http.StatusOK, res)
}

func (c *CustomerController) Show(ctx *gin.Context) {
	var errorOut []model.ErrorDetail
	id := ctx.Param("id")

	res, err := c.CustomerUseCase.GetById(ctx, id)
	if err != nil {
		errorOut = append(errorOut, err.Details)
		helper.ErrorResponse[any, []model.ErrorDetail](ctx, err.Code, errorOut)
		return
	}

	helper.SuccessResponse[*entity.Customer, any](ctx, http.StatusOK, res)
}

func (c *CustomerController) Update(ctx *gin.Context) {
	var errorOut []model.ErrorDetail
	var updateCustomerRequest model.UpdateCustomerRequest
	id := ctx.Param("id")

	if err := ctx.ShouldBindBodyWithJSON(&updateCustomerRequest); err != nil {
		helper.ValidationErrorResponse(ctx, err)
		return
	}

	res, err := c.CustomerUseCase.Update(ctx, id, &updateCustomerRequest)
	if err != nil {
		errorOut = append(errorOut, err.Details)
		helper.ErrorResponse[any, []model.ErrorDetail](ctx, err.Code, errorOut)
		return
	}

	helper.SuccessResponse[*entity.Customer, any](ctx, http.StatusOK, res)
}

func (c *CustomerController) Delete(ctx *gin.Context) {
	var errorOut []model.ErrorDetail
	id := ctx.Param("id")

	if err := c.CustomerUseCase.Delete(ctx, id); err != nil {
		errorOut = append(errorOut, err.Details)
		helper.ErrorResponse[any, []model.ErrorDetail](ctx, err.Code, errorOut)
		return
	}

	helper.SuccessResponse[any, any](ctx, http.StatusNoContent, nil)
}
