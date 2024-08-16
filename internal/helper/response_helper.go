package helper

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
	"waizly/internal/model"
	"waizly/pkg/apperrors"
)

func SuccessResponse[T any, U any](ctx *gin.Context, code int, data T) {
	ctx.JSON(code, model.ApiResponse[T, U]{Data: data})
}

func ErrorResponse[T any, U any](ctx *gin.Context, code int, details U) {
	ctx.AbortWithStatusJSON(code, model.ApiResponse[T, U]{Details: details})
}

func ErrorBindingResponse(ctx *gin.Context, validationErrors validator.ValidationErrors) {
	var responseErrors []model.ErrorDetail
	for _, fieldError := range validationErrors {
		field := fieldError.Field()
		responseErrors = append(responseErrors, model.ErrorDetail{
			Field:     field,
			ErrorCode: strings.ToUpper(fieldError.Tag()),
			Param:     fieldError.Param(),
			Message:   ParseFieldError(fieldError),
		})
	}
	ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, model.ApiResponse[any, []model.ErrorDetail]{Details: responseErrors})
}

func ValidationErrorResponse(ctx *gin.Context, err error) {
	var validationErrors validator.ValidationErrors

	if errors.As(err, &validationErrors) {
		ErrorBindingResponse(ctx, validationErrors)
	} else {
		var outErr []model.ErrorDetail
		outErr = append(outErr, model.ErrorDetail{
			Field:     "general",
			ErrorCode: string(apperrors.InvalidJson),
			Param:     "",
			Message:   err.Error(),
		})
		ErrorResponse[any, []model.ErrorDetail](ctx, http.StatusBadRequest, outErr)
	}
}
