package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"waizly/internal/helper"
	"waizly/internal/model"
	"waizly/internal/usecase"
)

type UserController struct {
	*usecase.UserUseCase
}

func NewUserController(userUseCase *usecase.UserUseCase) *UserController {
	return &UserController{UserUseCase: userUseCase}
}

func (u *UserController) GetUser(ctx *gin.Context) {
	userId, ok := ctx.Get("userId")
	if !ok {
		helper.ErrorResponse[any, string](ctx, http.StatusUnauthorized, "Unauthorized")
		ctx.Abort()
		return
	}

	res, err := u.UserUseCase.GetUser(ctx, &model.GetUserRequest{ID: userId.(uint)})
	if err != nil {
		helper.ErrorResponse[any, model.ErrorDetail](ctx, err.Code, err.Details)
		ctx.Abort()
		return
	}

	helper.SuccessResponse[*model.UserResponse, any](ctx, http.StatusOK, res)
}
