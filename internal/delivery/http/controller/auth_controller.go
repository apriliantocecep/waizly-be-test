package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"waizly/internal/helper"
	"waizly/internal/model"
	"waizly/internal/usecase"
)

type AuthController struct {
	UserUseCase *usecase.UserUseCase
}

func NewAuthController(userUseCase *usecase.UserUseCase) *AuthController {
	return &AuthController{UserUseCase: userUseCase}
}

func (c *AuthController) Login(ctx *gin.Context) {
	var errorOut []model.ErrorDetail
	var loginRequest model.LoginUserRequest

	if err := ctx.ShouldBindBodyWithJSON(&loginRequest); err != nil {
		helper.ValidationErrorResponse(ctx, err)
		return
	}

	res, err := c.UserUseCase.Login(ctx, &loginRequest)
	if err != nil {
		errorOut = append(errorOut, err.Details)
		helper.ErrorResponse[any, []model.ErrorDetail](ctx, err.Code, errorOut)
		return
	}

	helper.SuccessResponse[*model.UserResponse, any](ctx, http.StatusOK, res)
}

func (c *AuthController) Register(ctx *gin.Context) {
	var errorOut []model.ErrorDetail
	var registerRequest model.RegisterUserRequest

	if err := ctx.ShouldBindBodyWithJSON(&registerRequest); err != nil {
		helper.ValidationErrorResponse(ctx, err)
		return
	}

	res, err := c.UserUseCase.Register(ctx, &registerRequest)
	if err != nil {
		errorOut = append(errorOut, err.Details)
		helper.ErrorResponse[any, []model.ErrorDetail](ctx, err.Code, errorOut)
		return
	}

	helper.SuccessResponse[*model.UserResponse, any](ctx, http.StatusOK, res)
}
