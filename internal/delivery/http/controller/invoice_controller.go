package controller

import (
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"strconv"
	"waizly/internal/helper"
	"waizly/internal/model"
	"waizly/internal/usecase"
)

type InvoiceController struct {
	InvoiceUseCase *usecase.InvoiceUseCase
}

func NewInvoiceController(invoiceUseCase *usecase.InvoiceUseCase) *InvoiceController {
	return &InvoiceController{InvoiceUseCase: invoiceUseCase}
}

func (c *InvoiceController) Index(ctx *gin.Context) {
	var errorOut []model.ErrorDetail
	pageNumber := 1
	sizeNumber := 10
	totalNumber := 0

	page := ctx.Query("page")
	if pageInt, err := strconv.Atoi(page); err == nil {
		pageNumber = pageInt
	}

	size := ctx.Query("size")
	if sizeInt, err := strconv.Atoi(size); err == nil {
		sizeNumber = sizeInt
	}

	totalItems := ctx.Query("total_items")
	if totalInt, err := strconv.Atoi(totalItems); err == nil {
		totalNumber = totalInt
	}

	req := &model.SearchInvoiceRequest{
		ID:         ctx.Query("id"),
		IssueDate:  ctx.Query("issue_date"),
		DueDate:    ctx.Query("due_date"),
		Subject:    ctx.Query("subject"),
		Customer:   ctx.Query("customer"),
		Status:     ctx.Query("status"),
		TotalItems: totalNumber,
		Page:       pageNumber,
		Size:       sizeNumber,
	}

	res, total, err := c.InvoiceUseCase.Search(ctx, req)
	if err != nil {
		errorOut = append(errorOut, err.Details)
		helper.ErrorResponse[any, []model.ErrorDetail](ctx, err.Code, errorOut)
		return
	}

	paging := &model.PageMetadata{
		Page:      req.Page,
		Size:      req.Size,
		TotalItem: total,
		TotalPage: int64(math.Ceil(float64(total) / float64(req.Size))),
	}

	ctx.JSON(http.StatusOK, model.ApiPaginationResponse[[]model.InvoiceResponse]{
		Data:   res,
		Paging: paging,
	})
}

func (c *InvoiceController) Create(ctx *gin.Context) {
	var errorOut []model.ErrorDetail
	var createInvoiceRequest model.CreateInvoiceRequest

	if err := ctx.ShouldBindBodyWithJSON(&createInvoiceRequest); err != nil {
		helper.ValidationErrorResponse(ctx, err)
		return
	}

	res, err := c.InvoiceUseCase.Create(ctx, &createInvoiceRequest)
	if err != nil {
		errorOut = append(errorOut, err.Details)
		helper.ErrorResponse[any, []model.ErrorDetail](ctx, err.Code, errorOut)
		return
	}

	helper.SuccessResponse[*model.CreateInvoiceResponse, any](ctx, http.StatusOK, res)
}

func (c *InvoiceController) Show(ctx *gin.Context) {
	var errorOut []model.ErrorDetail
	id := ctx.Param("id")

	res, err := c.InvoiceUseCase.GetById(ctx, id)
	if err != nil {
		errorOut = append(errorOut, err.Details)
		helper.ErrorResponse[any, []model.ErrorDetail](ctx, err.Code, errorOut)
		return
	}

	helper.SuccessResponse[*model.InvoiceResponse, any](ctx, http.StatusOK, res)
}

func (c *InvoiceController) Delete(ctx *gin.Context) {
	var errorOut []model.ErrorDetail
	id := ctx.Param("id")

	if err := c.InvoiceUseCase.Delete(ctx, id); err != nil {
		errorOut = append(errorOut, err.Details)
		helper.ErrorResponse[any, []model.ErrorDetail](ctx, err.Code, errorOut)
		return
	}

	helper.SuccessResponse[any, any](ctx, http.StatusNoContent, nil)
}

func (c *InvoiceController) Update(ctx *gin.Context) {
	var errorOut []model.ErrorDetail
	var updateInvoiceRequest model.UpdateInvoiceRequest
	id := ctx.Param("id")

	if err := ctx.ShouldBindBodyWithJSON(&updateInvoiceRequest); err != nil {
		helper.ValidationErrorResponse(ctx, err)
		return
	}

	res, err := c.InvoiceUseCase.Update(ctx, id, &updateInvoiceRequest)
	if err != nil {
		errorOut = append(errorOut, err.Details)
		helper.ErrorResponse[any, []model.ErrorDetail](ctx, err.Code, errorOut)
		return
	}

	helper.SuccessResponse[*model.UpdateInvoiceResponse, any](ctx, http.StatusOK, res)
}
